package day3

import (
	"bufio"
	"fmt"
	"os"
)


func Part1() {
	f, err := os.Open("day3-input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	var score int
	s := bufio.NewScanner(f)
	for s.Scan() {
		ln := s.Text()
		x := scoreLine(ln)
		score += x

	}
	fmt.Println(score)
}

func scoreLine(ln string) int {
	r := findCommon(ln)
	x := scoreItem(r)
	return x
}

func findCommon(in string) rune {
	l, r := splitSack(in)
	return intersectSingle(l, r)
}

func splitSack(in string) (string, string) {
	i := len(in) / 2
	left := in[:i]
	right := in[i:]
	return left, right
}

func intersectSingle(l, r string) rune {
	for _, lb := range l {
		for _, rb := range r {
			if lb == rb {
				return lb
			}
		}
	}
	return rune(0)
}

func scoreItem(i rune) int {
	// Lowercase item types a through z have priorities 1 through 26.
	// Uppercase item types A through Z have priorities 27 through 52.
	int := int(i)
	if int >= 97 {
		return int - 96
	}

	return int - 38
}
