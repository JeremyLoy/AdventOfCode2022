package aoc

import (
	"embed"
	"strings"
	"testing"
)

//go:embed data
var testData embed.FS

func TestDay1CalorieCounting(t *testing.T) {
	exampleInput := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
	t.Run("part 1 example", func(t *testing.T) {

		largestElf, largestCount, err := GetLargestElf(strings.NewReader(exampleInput))
		if err != nil {
			t.Error(err)
		}
		if largestElf != 4 {
			t.Errorf("largest elf was %v, should have been 4", largestElf)
		}

		if largestCount != 24_000 {
			t.Errorf("largest count was %v, should have been 24000", largestElf)
		}
	})
	t.Run("part 1", func(t *testing.T) {
		day1, err := testData.Open("data/day1.txt")
		if err != nil {
			t.Fatal(err)
		}
		largestElf, largestCount, err := GetLargestElf(day1)
		if err != nil {
			t.Fatal(err)
		}
		if largestCount != 69_310 {
			t.Errorf("incorrect count - got %v", largestCount)
		}
		if largestElf != 178 {
			t.Errorf("incorrect elf - got %v", largestElf)
		}
	})
	t.Run("part 2 example", func(t *testing.T) {
		topThreeSum, err := GetSumThreeLargest(strings.NewReader(exampleInput))
		if err != nil {
			t.Fatal(err)
		}
		if topThreeSum != 45_000 {
			t.Errorf("incorrect sum - got %v", topThreeSum)
		}
	})
	t.Run("part 2", func(t *testing.T) {
		day1, err := testData.Open("data/day1.txt")
		if err != nil {
			t.Fatal(err)
		}
		topThreeSum, err := GetSumThreeLargest(day1)
		if err != nil {
			t.Fatal(err)
		}
		if topThreeSum != 206_104 {
			t.Errorf("incorrect sum - got %v", topThreeSum)
		}
	})
}
