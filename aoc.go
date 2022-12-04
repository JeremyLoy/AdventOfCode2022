package aoc

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

type Elf struct {
	number   int
	calories int
}

func GetElves(r io.Reader) ([]Elf, error) {
	scanner := bufio.NewScanner(r)
	var elves []Elf
	var currentElf Elf
	i := 1
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			elves = append(elves, currentElf)
			i++
			currentElf = Elf{number: i}
			continue
		}
		count, err := strconv.Atoi(text)
		if err != nil {
			return elves, fmt.Errorf("failed to convert text to number %v", err)
		}
		currentElf.calories += count
	}
	if scanner.Err() != nil {
		return elves, scanner.Err()
	}
	elves = append(elves, currentElf)
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].calories > elves[j].calories
	})
	return elves, nil
}

func GetLargestElf(r io.Reader) (int, int, error) {
	elves, err := GetElves(r)
	if err != nil {
		return 0, 0, err
	}
	return elves[0].number, elves[0].calories, nil
}

func GetSumThreeLargest(r io.Reader) (int, error) {
	elves, err := GetElves(r)
	if err != nil {
		return 0, err
	}
	return elves[0].calories + elves[1].calories + elves[2].calories, nil
}
