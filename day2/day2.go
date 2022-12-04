package day2

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Part1() {
	f, err := os.Open("day2-input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	var score int
	s := bufio.NewScanner(f)
	for s.Scan() {
		ln := s.Text()
		rawHands := strings.Split(ln, " ")
		oh, _ := parseOpponent(rawHands[0])
		uh, _ := parseYou(rawHands[1])

		w := winner(oh, uh)
		score += calcScore(uh, w)

	}
	fmt.Println(score)
}

func Part2() {
	f, err := os.Open("day2-input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	var score int
	s := bufio.NewScanner(f)
	for s.Scan() {
		ln := s.Text()
		rawHands := strings.Split(ln, " ")
		oh, _ := parseOpponent(rawHands[0])
		r := parseYouResult(rawHands[1])
		uh := calcYourHand(oh, r)
		w := winner(oh, uh)
		score += calcScore(uh, w)
	}
	fmt.Println(score)
}

type Hand int

const (
	Unknown  Hand = iota
	Rock     Hand = 1
	Paper    Hand = 2
	Scissors Hand = 3
)

type Result int

const (
	ResultUnknown Result = iota
	YouLose
	YouDraw
	YouWin
)

func parseOpponent(o string) (Hand, error) {
	switch o {
	case "A":
		return Rock, nil
	case "B":
		return Paper, nil
	case "C":
		return Scissors, nil
	}
	return Unknown, errors.New("Unknown")
}

func parseYou(u string) (Hand, error) {
	switch u {
	case "X":
		return Rock, nil
	case "Y":
		return Paper, nil
	case "Z":
		return Scissors, nil
	}
	return Unknown, errors.New("Unknown")
}

func parseYouResult(u string) Result {
	switch u {
	case "X":
		return YouLose
	case "Y":
		return YouDraw
	case "Z":
		return YouWin
	}
	return ResultUnknown
}

func calcYourHand(o Hand, r Result) Hand {
	if r == YouDraw {
		return o
	}

	if r == YouWin {
		switch {
		case o == Rock:
			return Paper
		case o == Paper:
			return Scissors
		default:
			return Rock
		}
	}

	if r == YouLose {
		switch {
		case o == Rock:
			return Scissors
		case o == Paper:
			return Rock
		default:
			return Paper
		}
	}
	return Unknown
}

func winner(oh Hand, uh Hand) Result {
	if oh == uh {
		return YouDraw
	}

	if oh == Rock {
		if uh == Scissors {
			return YouLose
		}
		return YouWin
	}

	if oh == Paper {
		if uh == Rock {
			return YouLose
		}
		return YouWin
	}

	if uh == Paper {
		return YouLose
	}

	return YouWin
}

func calcScore(u Hand, result Result) int {
	score := int(u)
	if result == YouDraw {
		score += 3
	} else if result == YouWin {
		score += 6
	}
	return score
}
