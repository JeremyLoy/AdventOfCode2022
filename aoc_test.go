package aoc

import (
	"embed"
	"io/fs"
	"strings"
	"testing"
)

//go:embed data
var testData embed.FS

func MustOpen(t *testing.T, name string) fs.File {
	t.Helper()
	file, err := testData.Open(name)
	if err != nil {
		t.Fatal(err)
	}
	return file
}

func TestDay1CalorieCounting(t *testing.T) {
	t.Parallel()
	exampleInput := "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000"
	t.Run("part 1 example", func(t *testing.T) {
		t.Parallel()
		elves, err := GetElves(strings.NewReader(exampleInput))
		if err != nil {
			t.Fatal(err)
		}
		largestElf := GetLargestElf(elves)

		if largestElf.Number != 4 {
			t.Errorf("largest elf was %v, should have been 4", largestElf.Number)
		}

		if largestElf.Calories != 24_000 {
			t.Errorf("largest count was %v, should have been 24000", largestElf.Calories)
		}
	})
	t.Run("part 1", func(t *testing.T) {
		t.Parallel()
		day1 := MustOpen(t, "data/day1.txt")
		elves, err := GetElves(day1)
		if err != nil {
			t.Fatal(err)
		}
		largestElf := GetLargestElf(elves)
		if largestElf.Calories != 69_310 {
			t.Errorf("incorrect count - got %v", largestElf.Calories)
		}
		if largestElf.Number != 178 {
			t.Errorf("incorrect elf - got %v", largestElf.Number)
		}
	})
	t.Run("part 2 example", func(t *testing.T) {
		t.Parallel()
		elves, err := GetElves(strings.NewReader(exampleInput))
		if err != nil {
			t.Fatal(err)
		}
		topThreeSum := SumThreeLargestElves(elves)

		if topThreeSum != 45_000 {
			t.Errorf("incorrect sum - got %v", topThreeSum)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		t.Parallel()
		day1 := MustOpen(t, "data/day1.txt")
		elves, err := GetElves(day1)
		if err != nil {
			t.Fatal(err)
		}
		topThreeSum := SumThreeLargestElves(elves)

		if topThreeSum != 206_104 {
			t.Errorf("incorrect sum - got %v", topThreeSum)
		}
	})
}

func TestDay2RockPaperScissors(t *testing.T) {
	t.Parallel()
	exampleStrategyGuide := "A Y\nB X\nC Z"
	t.Run("part 1 example", func(t *testing.T) {
		t.Parallel()
		strategyGuide, err := ParseStrategyGuide(strings.NewReader(exampleStrategyGuide))
		if err != nil {
			t.Fatal(err)
		}
		score := CalculateRPSScore(strategyGuide, ModeSelf)
		if score != 15 {
			t.Errorf("unexpected score %v", score)
		}
	})
	t.Run("part 1", func(t *testing.T) {
		t.Parallel()
		day2 := MustOpen(t, "data/day2.txt")
		strategyGuide, err := ParseStrategyGuide(day2)
		if err != nil {
			t.Fatal(err)
		}
		score := CalculateRPSScore(strategyGuide, ModeSelf)
		if score != 14_827 {
			t.Errorf("unexpected score %v", score)
		}
	})
	t.Run("part 2 example", func(t *testing.T) {
		t.Parallel()
		strategyGuide, err := ParseStrategyGuide(strings.NewReader(exampleStrategyGuide))
		if err != nil {
			t.Fatal(err)
		}
		score := CalculateRPSScore(strategyGuide, ModeOutcome)
		if score != 12 {
			t.Errorf("unexpected score %v", score)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		t.Parallel()
		day2 := MustOpen(t, "data/day2.txt")
		strategyGuide, err := ParseStrategyGuide(day2)
		if err != nil {
			t.Fatal(err)
		}
		score := CalculateRPSScore(strategyGuide, ModeOutcome)
		if score != 13_889 {
			t.Errorf("unexpected score %v", score)
		}
	})
}

func TestDay3RucksackReorganization(t *testing.T) {
	t.Parallel()
	exampleRucksacks := "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw"
	t.Run("part 1 example", func(t *testing.T) {
		t.Parallel()
		priority, err := SumPriority(strings.NewReader(exampleRucksacks))
		if err != nil {
			t.Fatal(err)
		}
		if priority != 157 {
			t.Errorf("unexpected priority %v", priority)
		}
	})
	t.Run("part 1", func(t *testing.T) {
		t.Parallel()
		day3 := MustOpen(t, "data/day3.txt")
		priority, err := SumPriority(day3)
		if err != nil {
			t.Fatal(err)
		}
		if priority != 8_176 {
			t.Errorf("unexpected priority %v", priority)
		}
	})
	t.Run("part 2 example", func(t *testing.T) {
		t.Parallel()
		priority, err := SumBadgePriority(strings.NewReader(exampleRucksacks))
		if err != nil {
			t.Fatal(err)
		}
		if priority != 70 {
			t.Errorf("unexpected priority %v", priority)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		t.Parallel()
		day3 := MustOpen(t, "data/day3.txt")
		priority, err := SumBadgePriority(day3)
		if err != nil {
			t.Fatal(err)
		}
		if priority != 2_689 {
			t.Errorf("unexpected priority %v", priority)
		}
	})
}

func TestDay4CampCleanup(t *testing.T) {
	t.Parallel()
	exampleAssignments := "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8"
	t.Run("part 1 example", func(t *testing.T) {
		t.Parallel()
		assignments, err := ParseAssignments(strings.NewReader(exampleAssignments))
		if err != nil {
			t.Fatal(err)
		}
		sum := SumFullyOverlaps(assignments)
		if sum != 2 {
			t.Errorf("unexpected sum of fully overlaps %v", sum)
		}
	})
	t.Run("part 1", func(t *testing.T) {
		t.Parallel()
		day4 := MustOpen(t, "data/day4.txt")
		assignments, err := ParseAssignments(day4)
		if err != nil {
			t.Fatal(err)
		}
		sum := SumFullyOverlaps(assignments)
		if sum != 466 {
			t.Errorf("unexpected sum of fully overlaps %v", sum)
		}
	})
	t.Run("part 2 example", func(t *testing.T) {
		t.Parallel()
		assignments, err := ParseAssignments(strings.NewReader(exampleAssignments))
		if err != nil {
			t.Fatal(err)
		}
		sum := SumOverlappingSections(assignments)
		if sum != 4 {
			t.Errorf("unexpected sum of overlapping sections %v", sum)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		t.Parallel()
		day4 := MustOpen(t, "data/day4.txt")
		assignments, err := ParseAssignments(day4)
		if err != nil {
			t.Fatal(err)
		}
		sum := SumOverlappingSections(assignments)
		if sum != 865 {
			t.Errorf("unexpected sum of overlapping sections %v", sum)
		}
	})
}

func TestDay5SupplyStacks(t *testing.T) {
	t.Parallel()
	t.Run("part 1 example", func(t *testing.T) {
		t.Parallel()
		day5Example := MustOpen(t, "data/day5Example.txt")
		stacks, steps, err := ParseStacksAndSteps(day5Example)
		if err != nil {
			t.Fatal(err)
		}
		stacks = ProcessSteps(stacks, steps)
		message := SumTopOfStacks(stacks)
		if message != "CMZ" {
			t.Errorf("unexpected message %v", message)
		}
	})
	t.Run("part 1", func(t *testing.T) {
		t.Parallel()
		day5 := MustOpen(t, "data/day5.txt")
		stacks, steps, err := ParseStacksAndSteps(day5)
		if err != nil {
			t.Fatal(err)
		}
		stacks = ProcessSteps(stacks, steps)
		message := SumTopOfStacks(stacks)
		if message != "DHBJQJCCW" {
			t.Errorf("unexpected message %v", message)
		}
	})
	t.Run("part 2 example", func(t *testing.T) {
		t.Parallel()
		day5Example := MustOpen(t, "data/day5Example.txt")
		stacks, steps, err := ParseStacksAndSteps(day5Example)
		if err != nil {
			t.Fatal(err)
		}
		stacks = ProcessSteps9001(stacks, steps)
		message := SumTopOfStacks(stacks)
		if message != "MCD" {
			t.Errorf("unexpected message %v", message)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		t.Parallel()
		day5 := MustOpen(t, "data/day5.txt")
		stacks, steps, err := ParseStacksAndSteps(day5)
		if err != nil {
			t.Fatal(err)
		}
		stacks = ProcessSteps9001(stacks, steps)
		message := SumTopOfStacks(stacks)
		if message != "WJVRLSJJT" {
			t.Errorf("unexpected message %v", message)
		}
	})
}

func TestDay6TuningTrouble(t *testing.T) {
	t.Parallel()
	t.Run("part 1 example", func(t *testing.T) {
		t.Parallel()
		if start := StartOfPacket(strings.NewReader("mjqjpqmgbljsphdztnvjfqwrcgsmlb")); start != 7 {
			t.Errorf("unexpected first start %v", start)
		}
		if start := StartOfPacket(strings.NewReader("bvwbjplbgvbhsrlpgdmjqwftvncz")); start != 5 {
			t.Errorf("unexpected first start %v", start)
		}
		if start := StartOfPacket(strings.NewReader("nppdvjthqldpwncqszvftbrmjlhg")); start != 6 {
			t.Errorf("unexpected first start %v", start)
		}
		if start := StartOfPacket(strings.NewReader("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")); start != 10 {
			t.Errorf("unexpected first start %v", start)
		}
		if start := StartOfPacket(strings.NewReader("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")); start != 11 {
			t.Errorf("unexpected first start %v", start)
		}
	})
	t.Run("part 1", func(t *testing.T) {
		t.Parallel()
		day6 := MustOpen(t, "data/day6.txt")
		if start := StartOfPacket(day6); start != 1042 {
			t.Errorf("unexpected start %v", start)
		}
	})
	t.Run("part 2 example", func(t *testing.T) {
		t.Parallel()
		if start := StartOfMessage(strings.NewReader("mjqjpqmgbljsphdztnvjfqwrcgsmlb")); start != 19 {
			t.Errorf("unexpected first start %v", start)
		}
		if start := StartOfMessage(strings.NewReader("bvwbjplbgvbhsrlpgdmjqwftvncz")); start != 23 {
			t.Errorf("unexpected first start %v", start)
		}
		if start := StartOfMessage(strings.NewReader("nppdvjthqldpwncqszvftbrmjlhg")); start != 23 {
			t.Errorf("unexpected first start %v", start)
		}
		if start := StartOfMessage(strings.NewReader("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")); start != 29 {
			t.Errorf("unexpected first start %v", start)
		}
		if start := StartOfMessage(strings.NewReader("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")); start != 26 {
			t.Errorf("unexpected first start %v", start)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		t.Parallel()
		day6 := MustOpen(t, "data/day6.txt")
		if start := StartOfMessage(day6); start != 2980 {
			t.Errorf("unexpected start %v", start)
		}
	})
}

func TestDay7NoSpaceLeftOnDevice(t *testing.T) {
	t.Parallel()
	t.Run("part 1 example", func(t *testing.T) {
		t.Parallel()
		day7example := MustOpen(t, "data/day7example.txt")
		root, err := ParseFS(day7example)
		if err != nil {
			t.Fatal(err)
		}
		dirs := root.GetDirs()
		totalSize := SumDirSize(dirs)
		if totalSize != 95_437 {
			t.Errorf("unexpected total size %v", totalSize)
		}
	})
	t.Run("part 1", func(t *testing.T) {
		t.Parallel()
		day7 := MustOpen(t, "data/day7.txt")
		root, err := ParseFS(day7)
		if err != nil {
			t.Fatal(err)
		}
		dirs := root.GetDirs()
		totalSize := SumDirSize(dirs)
		if totalSize != 1_297_683 {
			t.Errorf("unexpected total size %v", totalSize)
		}
	})
	t.Run("part 2 example", func(t *testing.T) {
		t.Parallel()
		day7example := MustOpen(t, "data/day7example.txt")
		root, err := ParseFS(day7example)
		if err != nil {
			t.Fatal(err)
		}
		deletedSize := SmallestDirToDelete(root)
		if deletedSize != 24_933_642 {
			t.Errorf("unexpected deleted size of %v", deletedSize)
		}

	})
	t.Run("part 2", func(t *testing.T) {
		t.Parallel()
		day7 := MustOpen(t, "data/day7.txt")
		root, err := ParseFS(day7)
		if err != nil {
			t.Fatal(err)
		}
		deletedSize := SmallestDirToDelete(root)
		if deletedSize != 5_756_764 {
			t.Errorf("unexpected deleted size of %v", deletedSize)
		}
	})
}
