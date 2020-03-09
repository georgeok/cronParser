package cronParser

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// bounds provides a range of acceptable values (plus a map of name to value).
type bounds struct {
	min, max int
	names    map[string]int
}

// Any represents the cron "any value"
const (
	anyChar   = "*"
	rangeChar = "-"
	stepChar  = "/"
	listChar  = ","
)

// The bounds for each field.
var (
	Minutes     = bounds{0, 59, nil}
	Hours       = bounds{0, 23, nil}
	DaysOfMonth = bounds{1, 31, nil}
	Months      = bounds{1, 12, map[string]int{
		"jan": 1,
		"feb": 2,
		"mar": 3,
		"apr": 4,
		"may": 5,
		"jun": 6,
		"jul": 7,
		"aug": 8,
		"sep": 9,
		"oct": 10,
		"nov": 11,
		"dec": 12,
	}}
	DaysOfWeek = bounds{0, 6, map[string]int{
		"sun": 0,
		"mon": 1,
		"tue": 2,
		"wed": 3,
		"thu": 4,
		"fri": 5,
		"sat": 6,
	}}
)

// Parse contains the main logic for parsing the cron commands and return the formatted output
func Parse(input string, bounds bounds) string {
	input = replaceNameWithIntegers(input, bounds)
	listItems := lists(input)
	out, err := cronRange(listItems, bounds)
	if err != nil {
		return "Wrong input"
	}
	sort.Ints(out)
	str := fmt.Sprint(out)
	return str[1 : len(str)-1]
}

func replaceNameWithIntegers(input string, bounds bounds) string {
	input = strings.ToLower(input)
	if bounds.names != nil {
		for k, v := range bounds.names {
			input = strings.ReplaceAll(input, k, strconv.Itoa(v))
		}
	}
	return input
}

func lists(input string) []string {
	return strings.Split(input, listChar)
}

func steps(input string) (string, int) {
	if !strings.Contains(input, stepChar) {
		return input, 1
	}
	stepStr := strings.Split(input, stepChar)
	step, err := strconv.Atoi(stepStr[1])
	if err != nil {
		return input, 1
	}
	return stepStr[0], step
}

func cronRange(lists []string, bounds bounds) ([]int, error) {
	ranges := make([]int, 0)
	for _, val := range lists {
		from := bounds.min
		to := bounds.max
		stepVal := 1
		var err error

		switch {
		case step(val):
			stepStr := strings.Split(val, stepChar)
			stepVal, err = strconv.Atoi(stepStr[1])
			if err != nil {
				break
			}

			if rng(stepStr[0]) {
				fromTo := strings.Split(stepStr[0], rangeChar)
				from, err = strconv.Atoi(fromTo[0])
				if err != nil {
					break
				}
				to, err = strconv.Atoi(fromTo[1])
				if err != nil {
					break
				}
			} else if any(stepStr[0]) {
				from = bounds.min
				to = bounds.max
			} else {
				from, err = strconv.Atoi(stepStr[0])
				if err != nil {
					break
				}
				to = bounds.max
			}
		case rng(val):
			fromTo := strings.Split(val, rangeChar)
			from, err = strconv.Atoi(fromTo[0])
			if err != nil {
				break
			}
			to, err = strconv.Atoi(fromTo[1])
			if err != nil {
				break
			}
		case any(val):
			from = bounds.min
			to = bounds.max
		default:
			from, err = strconv.Atoi(val)
			if err != nil {
				break
			}
			to = from
		}
		if err != nil || from < bounds.min || to > bounds.max || from > to || stepVal < 1 {
			return []int{}, err
		}
		for i := from; i <= to; i += stepVal {
			ranges = append(ranges, i)
		}
	}
	return ranges, nil
}

func any(input string) bool {
	return input == anyChar
}

func rng(input string) bool {
	return strings.Contains(input, rangeChar)
}

func step(input string) bool {
	return strings.Contains(input, stepChar)
}
