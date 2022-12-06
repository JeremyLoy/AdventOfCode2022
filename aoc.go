package aoc

import (
	"bufio"
	"errors"
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
	UnknownOutcome Outcome = -1
	Lose           Outcome = 0
	Draw           Outcome = 3
	Win            Outcome = 6
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
		var outcome Outcome
		switch r.Self - r.Opponent {
		case 0:
			outcome = Draw
		case 1, -2:
			outcome = Win
		case -1, 2:
			outcome = Lose
		}
		return int(r.Self) + int(outcome)
	case ModeOutcome:
		var choice Shape
		switch r.Outcome {
		case Lose:
			switch r.Opponent {
			case Scissors:
				choice = Paper
			case Rock:
				choice = Scissors
			case Paper:
				choice = Rock
			}
		case Draw:
			choice = r.Opponent
		case Win:
			choice = r.Opponent%3 + 1
		}
		return int(choice) + int(r.Outcome)
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

type Rucksack string

var priorityAlphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Priority(rucksack string) int {
	mid := len(rucksack) / 2
	before, after := rucksack[mid:], rucksack[:mid]
	// alternatively, call BadgePriority with [before,after]
	r := before[strings.IndexAny(before, after)]
	return strings.IndexByte(string(priorityAlphabet), r) + 1
}

func BadgePriority(rucksacks []string) int {
	seen := make(map[rune]map[int]struct{})
	for i, rucksack := range rucksacks {
		for _, r := range rucksack {
			if _, ok := seen[r]; !ok {
				seen[r] = make(map[int]struct{})
			}
			seen[r][i] = struct{}{}
		}
	}
	// should only have one seen in all
	for r, v := range seen {
		if len(v) == len(rucksacks) {
			return strings.IndexRune(string(priorityAlphabet), r) + 1
		}
	}
	return -1
}

func SumPriority(r io.Reader) (int, error) {
	var sum int
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		sum += Priority(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return sum, nil
}

func SumBadgePriority(r io.Reader) (int, error) {
	var sum int
	scanner := bufio.NewScanner(r)
	var rucksacks []string
	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())
		if len(rucksacks) == 3 {
			sum += BadgePriority(rucksacks)
			rucksacks = nil
			continue
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return sum, nil
}

type Assignment struct {
	Start int
	End   int
}

func (a Assignment) FullyContains(b Assignment) bool {
	return b.Start <= a.End && b.Start >= a.Start && b.End <= a.End && b.End >= a.Start
}

type AssignmentPair struct {
	Left  Assignment
	Right Assignment
}

func (a AssignmentPair) String() string {
	return fmt.Sprintf("%v-%v,%v-%v", a.Left.Start, a.Left.End, a.Right.Start, a.Right.End)
}

func (p AssignmentPair) FullyContains() bool {
	return p.Left.FullyContains(p.Right) || p.Right.FullyContains(p.Left)
}

// OverlappingSections returns the count of overlapping sections of the pair. Its not actually used in AoC, but
// I wrote it by accident after misreading the prompt. I kept it as an implementation detail of [AssignmentPair.HasOverlappingSections]
func (p AssignmentPair) OverlappingSections() int {
	s := p.String()
	_ = s
	if p.Left.FullyContains(p.Right) {
		return p.Right.End - p.Right.Start + 1
	}
	if p.Right.FullyContains(p.Left) {
		return p.Left.End - p.Left.Start + 1
	}
	if p.Left.End >= p.Right.Start && p.Left.End <= p.Right.End {
		return p.Left.End - p.Right.Start + 1
	}
	if p.Right.End >= p.Left.Start && p.Right.End <= p.Left.End {
		return p.Right.End - p.Left.Start + 1
	}
	return 0
}

func (p AssignmentPair) HasOverlappingSections() bool {
	return p.OverlappingSections() > 0
}

func ParseAssignments(r io.Reader) ([]AssignmentPair, error) {
	scanner := bufio.NewScanner(r)
	var assignmentPairs []AssignmentPair
	for scanner.Scan() {
		var pair AssignmentPair
		count, err := fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &(pair.Left.Start), &(pair.Left.End), &(pair.Right.Start), &(pair.Right.End))
		if count != 4 || err != nil {
			return nil, fmt.Errorf("failed to scan assignment pair count: %v, err: %v", count, err)
		}
		assignmentPairs = append(assignmentPairs, pair)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return assignmentPairs, nil
}

func SumFullyOverlaps(assignments []AssignmentPair) int {
	var sum int
	for _, pair := range assignments {
		if pair.FullyContains() {
			sum++
		}
	}
	return sum
}

func SumOverlappingSections(assignments []AssignmentPair) int {
	var sum int
	for _, pair := range assignments {
		if pair.HasOverlappingSections() {
			sum++
		}
	}
	return sum
}

type Step struct {
	Amount, From, To int
}

type Stack []string

func (s *Stack) Push(st ...string) {
	*s = append(*s, st...)
}

func (s *Stack) Peek() string {
	i := len(*s) - 1
	return (*s)[i]
}

func (s *Stack) Pop() string {
	i := len(*s) - 1
	elem := (*s)[i]
	*s = (*s)[:i]
	return elem
}

func (s *Stack) PopN(amount int) []string {
	i := len(*s) - amount
	elems := (*s)[i:]
	*s = (*s)[:i]
	return elems
}

func (s *Stack) Len() int {
	return len(*s)
}

func ParseStacksAndSteps(r io.Reader) ([]*Stack, []Step, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, nil, err
	}

	stackHalf, stepHalf, found := strings.Cut(string(b), "\n\n")
	if !found {
		return nil, nil, errors.New("failed to split stacks and steps")
	}

	var stacks []*Stack
	var steps []Step

	stacksStrings := strings.Split(stackHalf, "\n")
	// count the number of integers in the final row
	width := len(strings.Fields(stacksStrings[len(stacksStrings)-1]))
	stacks = make([]*Stack, width)
	for i := range stacks {
		stacks[i] = new(Stack)
	}

	// reversed, skip the integer row as well
	for i := len(stacksStrings) - 2; i >= 0; i-- {
		stackString := stacksStrings[i]
		var current string
		for i := 0; i < width; i++ {
			// [Z]\w is 4 characters, but the final one is only 3. Avoid a nil pointer error
			if len(stackString) < 4 {
				current = stackString
			} else {
				current, stackString = stackString[:4], stackString[4:]
			}
			var crate string
			n, err := fmt.Sscanf(current, "[%1s]", &crate)
			if n != 1 || err != nil {
				continue
			}
			stacks[i].Push(crate)
		}
	}

	stepsStrings := strings.Split(stepHalf, "\n")
	for _, stepString := range stepsStrings {
		var step Step
		fmt.Sscanf(stepString, "move %d from %d to %d", &step.Amount, &step.From, &step.To)
		// to zero index it all
		step.From--
		step.To--
		steps = append(steps, step)
	}

	return stacks, steps, nil
}

func ProcessSteps(stacks []*Stack, steps []Step) []*Stack {
	for _, step := range steps {
		for i := 0; i < step.Amount; i++ {
			from := stacks[step.From]
			to := stacks[step.To]
			crate := from.Pop()
			to.Push(crate)
		}
	}
	return stacks
}
func ProcessSteps9001(stacks []*Stack, steps []Step) []*Stack {
	for _, step := range steps {
		from := stacks[step.From]
		to := stacks[step.To]
		crates := from.PopN(step.Amount)
		to.Push(crates...)
	}
	return stacks
}
func SumTopOfStacks(stacks []*Stack) string {
	var sum []string
	for _, stack := range stacks {
		sum = append(sum, stack.Peek())
	}
	return strings.Join(sum, "")
}
