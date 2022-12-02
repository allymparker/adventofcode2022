package day1

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
)

func Main() {
	f, err := os.Open("day1-input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var maxElvesCals [3]int

	elf := 0
	elfCals := 0

	s := bufio.NewScanner(f)
	for s.Scan() {
		ln := s.Text()
		if ln == "" {
			fmt.Printf("Total Cals for elf %d: %d\n", elf, elfCals)
			for i := 0; i < 3; i++ {
				if elfCals > maxElvesCals[i] {
					maxElvesCals[i] = elfCals
					sort.Ints(maxElvesCals[:])
					break
				}
			}
			elf++
			elfCals = 0
		}
		cals, _ := strconv.Atoi(ln)
		elfCals += cals
	}

	fmt.Println("\nTop 3")
	var totalMaxCals int
	for i := 0; i < 3; i++ {
		totalMaxCals+=maxElvesCals[i]
		fmt.Println(maxElvesCals[i])
	}
	fmt.Printf("\nTotal Top 3 Cals %d", totalMaxCals)
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
