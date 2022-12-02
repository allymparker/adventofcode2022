package day2

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Main() {
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

		score += scoreHand(rawHands[0], rawHands[1])

	}
	fmt.Println(score)
}

type Hand int

const (
	Unknown  Hand = iota
	Rock          = 1
	Paper         = 2
	Scissors      = 3
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

func scoreHand(o string, u string) int {
	oh, err := parseOpponent(o)
	if err != nil {
		return 0
	}

	uh, err := parseYou(u)
	if err != nil {
		return 0
	}

	w := winner(oh, uh)

	return score(uh, w)
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

func score(u Hand, result Result) int {
	score := int(u)
	if result == YouDraw {
		score += 3
	} else if result == YouWin {
		score += 6
	}
	return score
}
