package day3

import (
	"testing"
)

func Test_splitSack(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{"Split", args{"AB"}, "A", "B"},
		{"Split", args{"AABB"}, "AA", "BB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitSack(tt.args.in)
			if got != tt.want {
				t.Errorf("splitSack() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("splitSack() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_findCommon(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"Common", args{"AA"}, 'A'},
		{"Common", args{"bACdeA"}, 'A'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCommon(tt.args.in); got != tt.want {
				t.Errorf("findCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intersectSingle(t *testing.T) {
	type args struct {
		l string
		r string
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"no", args{"abc", "def"}, rune(0)},
		{"yes", args{"abc", "deb"}, 'b'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectSingle(tt.args.l, tt.args.r); got != tt.want {
				t.Errorf("intersectSingle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scoreItem(t *testing.T) {
	type args struct {
		i rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"a", args{'a'}, 1},
		{"z", args{'z'}, 26},
		{"A", args{'A'}, 27},
		{"Z", args{'Z'}, 52},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreItem(tt.args.i); got != tt.want {
				t.Errorf("scoreItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scoreLine(t *testing.T) {
	type args struct {
		ln string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"vJrwpWtwJgWrhcsFMMfFFhFp"}, 16},
		{"1", args{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"}, 38},
		{"1", args{"PmmdzqPrVvPwwTWBwg"}, 42},
		{"1", args{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn"}, 22},
		{"1", args{"ttgJtRGJQctTZtZT"}, 20},
		{"1", args{"CrZsJsPPZsGzwwsLwLmpwMDw"}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scoreLine(tt.args.ln); got != tt.want {
				t.Errorf("scoreLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
