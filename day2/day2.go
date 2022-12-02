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

	w, _ := handWinner(oh, uh)

	return score(uh, w)
}

func winner(o string, u string) (int, error) {
	oh, err := parseOpponent(o)
	if err != nil {
		return -1, errors.New("couldn't parse o")
	}

	uh, err := parseYou(u)
	if err != nil {
		return -1, errors.New("couldn't parse u")
	}

	return handWinner(oh, uh)
}

func handWinner(oh Hand, uh Hand) (int, error) {
	if oh == uh {
		return 0, nil
	}

	if oh == Rock {
		if uh == Scissors {
			return 1, nil
		}
		return 2, nil
	}

	if oh == Paper {
		if uh == Rock {
			return 1, nil
		}
		return 2, nil
	}

	if uh == Paper {
		return 1, nil
	}

	return 2, nil
}

func score(u Hand, result int) int {
	score := int(u)
	if result == 0 {
		score += 3
	} else if result == 2 {
		score += 6
	}
	return score
}
