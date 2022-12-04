package aoc

import (
	"embed"
	"strings"
	"testing"
)

//go:embed data
var testData embed.FS

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

		if largestElf.number != 4 {
			t.Errorf("largest elf was %v, should have been 4", largestElf.number)
		}

		if largestElf.calories != 24_000 {
			t.Errorf("largest count was %v, should have been 24000", largestElf.calories)
		}
	})
	t.Run("part 1", func(t *testing.T) {
		t.Parallel()
		day1, err := testData.Open("data/day1.txt")
		if err != nil {
			t.Fatal(err)
		}
		elves, err := GetElves(day1)
		if err != nil {
			t.Fatal(err)
		}
		largestElf := GetLargestElf(elves)
		if largestElf.calories != 69_310 {
			t.Errorf("incorrect count - got %v", largestElf.calories)
		}
		if largestElf.number != 178 {
			t.Errorf("incorrect elf - got %v", largestElf.number)
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
		day1, err := testData.Open("data/day1.txt")
		if err != nil {
			t.Fatal(err)
		}
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
