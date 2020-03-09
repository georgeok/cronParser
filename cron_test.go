package cronParser

import (
	"reflect"
	"testing"
)

func Test_minute(t *testing.T) {
	tests := []struct {
		name       string
		wantOutput string
	}{
		{"*", "0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 " +
			"21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 " +
			"45 46 47 48 49 50 51 52 53 54 55 56 57 58 59"},
		{"*/15", "0 15 30 45"},
		{"1-5", "1 2 3 4 5"},
		{"1,2,3", "1 2 3"},
		{"1", "1"},
		{"1-3,10-11", "1 2 3 10 11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.name, Minutes); !reflect.DeepEqual(got, tt.wantOutput) {
				t.Errorf("cronRange() = %v, want %v", got, tt.wantOutput)
			}
		})
	}

}

func Test_hour(t *testing.T) {
	tests := []struct {
		name       string
		wantOutput string
	}{
		{"*", "0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23"},
		{"20/2", "20 22"},
		{"1-5", "1 2 3 4 5"},
		{"1,2,3", "1 2 3"},
		{"1", "1"},
		{"1-3,10-11", "1 2 3 10 11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.name, Hours); !reflect.DeepEqual(got, tt.wantOutput) {
				t.Errorf("cronRange() = %v, want %v", got, tt.wantOutput)
			}
		})
	}

}

func Test_dayOfMonth(t *testing.T) {
	tests := []struct {
		name       string
		wantOutput string
	}{
		{"*", "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31"},
		{"20/3", "20 23 26 29"},
		{"1-5", "1 2 3 4 5"},
		{"1,2,3", "1 2 3"},
		{"1", "1"},
		{"1-3,10-11", "1 2 3 10 11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.name, DaysOfMonth); !reflect.DeepEqual(got, tt.wantOutput) {
				t.Errorf("cronRange() = %v, want %v", got, tt.wantOutput)
			}
		})
	}

}
func Test_month(t *testing.T) {
	tests := []struct {
		name       string
		wantOutput string
	}{
		{"*", "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31"},
		{"1/10", "1 11 21 31"},
		{"1-5", "1 2 3 4 5"},
		{"1,2,3", "1 2 3"},
		{"1", "1"},
		{"1-3,10-11", "1 2 3 10 11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.name, DaysOfMonth); !reflect.DeepEqual(got, tt.wantOutput) {
				t.Errorf("cronRange() = %v, want %v", got, tt.wantOutput)
			}
		})
	}

}
func Test_dayOfWeek(t *testing.T) {
	tests := []struct {
		name       string
		wantOutput string
	}{
		{"*", "0 1 2 3 4 5 6"},
		{"0/6", "0 6"},
		{"1-5", "1 2 3 4 5"},
		{"1,2,3", "1 2 3"},
		{"1", "1"},
		{"1-3,10-11", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.name, DaysOfWeek); !reflect.DeepEqual(got, tt.wantOutput) {
				t.Errorf("cronRange() = %v, want %v", got, tt.wantOutput)
			}
		})
	}
}

func Test_steps(t *testing.T) {
	tests := []struct {
		name  string
		want  string
		want1 int
	}{
		{"*/15", "*", 15},
		{"*/1", "*", 1},
		{"*", "*", 1},
		{"*/-1", "*", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := steps(tt.name)
			if got != tt.want {
				t.Errorf("steps() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("steps() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_lists(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{"15", []string{"15"}},
		{"1-10", []string{"1-10"}},
		{"1-10,11-20", []string{"1-10", "11-20"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lists(tt.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cronRange(t *testing.T) {
	tests := []struct {
		name   string
		lists  []string
		bounds bounds
		want   []int
	}{
		{"0-3", []string{"0-3"}, Months, []int{}},
		{"1-3", []string{"1-3"}, Minutes, []int{1, 2, 3}},
		{"1-3,11-20", []string{"1-3", "11-12"}, Minutes, []int{1, 2, 3, 11, 12}},
		{"1-3,59-80", []string{"1-3", "59-80"}, Minutes, []int{}},
		{"*", []string{"*"}, DaysOfWeek, []int{0, 1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := cronRange(tt.lists, tt.bounds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cronRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name   string
		bounds bounds
		want   string
	}{
		{"MON-TUE", DaysOfWeek, "1 2"},
		{"JAN-FEB", Months, "1 2"},
		{"1-3,11-12", Minutes, "1 2 3 11 12"},
		{"1-2,11", Minutes, "1 2 11"},
		{"1,20/2", Hours, "1 20 22"},
		{"1,2,20/2", Hours, "1 2 20 22"},
		{"1,2,20/aaa", Hours, "Wrong input"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.name, tt.bounds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cronRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
