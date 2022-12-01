package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("day1-input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	elf := 0
	elfCals := 0
	maxCalsElf := 0
	maxCals := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		ln := s.Text()
		if ln == "" {
			if elfCals >= maxCals {
				fmt.Printf("Elf %d has more calories than elf %d: %d > %d\n", elf, maxCalsElf, elfCals, maxCals)

				maxCals = elfCals
				maxCalsElf = elf
			}
			elf++
			elfCals = 0
		}
		cals, _ := strconv.Atoi(ln)
		elfCals += cals
	}

	fmt.Printf("\nElf %d has most calories %d", maxCalsElf, maxCals)
}

func downloadInput() (err error) {
	f, err := os.Create("day1-input")
	if err != nil {
		return err
	}
	defer f.Close()

	r, err := http.Get("https://adventofcode.com/2022/day/1/input")
	if err != nil {
		return err
	}
	defer r.Body.Close()

	_, err = io.Copy(f, r.Body)
	if err != nil {
		return err
	}

	return nil
}
