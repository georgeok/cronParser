package main

import (
	"fmt"
	"os"

	"github.com/georgeok/cronParser"
)

func main() {
	if len(os.Args) != 7 {
		fmt.Println("Please add exactly 7 arguments")
		fmt.Println("Seen", len(os.Args), os.Args)
		os.Exit(-1)
	}

	minute := cronParser.Parse(os.Args[1], cronParser.Minutes)
	hour := cronParser.Parse(os.Args[2], cronParser.Hours)
	dayOfMonth := cronParser.Parse(os.Args[3], cronParser.DaysOfMonth)
	month := cronParser.Parse(os.Args[4], cronParser.Months)
	dayOfWeek := cronParser.Parse(os.Args[5], cronParser.DaysOfWeek)
	cmd := os.Args[6]

	format := "%-14s%s\n"
	fmt.Printf(format, "minute", minute)
	fmt.Printf(format, "hour", hour)
	fmt.Printf(format, "day of month", dayOfMonth)
	fmt.Printf(format, "month", month)
	fmt.Printf(format, "day of week", dayOfWeek)
	fmt.Printf(format, "command", cmd)
}
