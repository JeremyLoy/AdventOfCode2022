package aoc

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	Number   int
	Calories int
}

// GetElves returns a list of all elves along with the number of nalories they are holding, sorted descending by calorie count
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
			currentElf = Elf{Number: i}
			continue
		}
		count, err := strconv.Atoi(text)
		if err != nil {
			return elves, fmt.Errorf("failed to convert text to Number %v", err)
		}
		currentElf.Calories += count
	}
	if scanner.Err() != nil {
		return elves, scanner.Err()
	}
	elves = append(elves, currentElf)
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].Calories > elves[j].Calories
	})
	return elves, nil
}

func GetLargestElf(elves []Elf) Elf {
	return elves[0]
}
func SumThreeLargestElves(elves []Elf) int {
	return elves[0].Calories + elves[1].Calories + elves[2].Calories
}

type Shape int

const (
	UnknownShape = iota
	Rock
	Paper
	Scissors
)

func NewShape(s string) Shape {
	switch s {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	default:
		return UnknownShape
	}

}

type Outcome int

const (
	UnknownOutcome = iota
	Lose
	Draw
	Win
)

func NewOutcome(s string) Outcome {
	switch s {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	default:
		return UnknownOutcome
	}
}

type RPSRound struct {
	Opponent Shape
	Self     Shape
	Outcome  Outcome
}

func (r RPSRound) Score(mode StrategyGuideMode) int {
	switch mode {
	case ModeSelf:
		var outcome int
		switch r.Self - r.Opponent {
		case 0:
			outcome = 3
		case 1, -2:
			outcome = 6
		case -1, 2:
			outcome = 0
		}
		return int(r.Self) + outcome
	case ModeOutcome:
		var outcome int
		var choice Shape
		switch r.Outcome {
		case Lose:
			outcome = 0
			switch r.Opponent {
			case Scissors:
				choice = Paper
			case Rock:
				choice = Scissors
			case Paper:
				choice = Rock
			}
		case Draw:
			outcome = 3
			choice = r.Opponent
		case Win:
			outcome = 6
			switch r.Opponent {
			case Scissors:
				choice = Rock
			case Rock:
				choice = Paper
			case Paper:
				choice = Scissors
			}
		}
		return int(choice) + outcome
	default:
		return 0
	}
}

type StrategyGuideMode int

const (
	ModeSelf StrategyGuideMode = iota + 1
	ModeOutcome
)

func ParseStrategyGuide(r io.Reader) ([]RPSRound, error) {
	scanner := bufio.NewScanner(r)
	var rounds []RPSRound
	for scanner.Scan() {
		var round RPSRound
		before, after, found := strings.Cut(scanner.Text(), " ")
		if !found {
			return nil, fmt.Errorf("error parsing guide, '%v'", scanner.Text())
		}
		round.Opponent = NewShape(before)
		round.Self = NewShape(after)
		round.Outcome = NewOutcome(after)
		rounds = append(rounds, round)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return rounds, nil
}

func CalculateRPSScore(strategyGuide []RPSRound, mode StrategyGuideMode) int {
	var total int
	for _, round := range strategyGuide {
		total += round.Score(mode)
	}
	return total
}
