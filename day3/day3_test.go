package day3_test

import (
	"bufio"
	"github.com/xorilog/aoc2022/day3"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var example = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

func TestRucksackContent(t *testing.T) {
	rucksacks := day3.FillRucksacks(example)
	// content := rucksacks[0].Content()
	assert.Equal(t, "vJrwpWtwJgWrhcsFMMfFFhFp", rucksacks[0].Content())
	assert.Equal(t, "vJrwpWtwJgWr", rucksacks[0].One())
	assert.Equal(t, "hcsFMMfFFhFp", rucksacks[0].Two())
	assert.Equal(t, []string{"p"}, rucksacks[0].Intersection())
	assert.Equal(t, "jqHRNqRjqzjGDLGL", rucksacks[1].One())
	assert.Equal(t, "rsFMfFZSrLrFZsSL", rucksacks[1].Two())
	assert.Equal(t, []string{"L"}, rucksacks[1].Intersection())
	assert.Equal(t, "PmmdzqPrV", rucksacks[2].One())
	assert.Equal(t, "vPwwTWBwg", rucksacks[2].Two())
	assert.Equal(t, []string{"P"}, rucksacks[2].Intersection())
	assert.Equal(t, []string{"v"}, rucksacks[3].Intersection())
	assert.Equal(t, []string{"t"}, rucksacks[4].Intersection())
	assert.Equal(t, []string{"s"}, rucksacks[5].Intersection())
	assert.Equal(t, 16, day3.Priority(rucksacks[0].Intersection()[0]))
	assert.Equal(t, 38, day3.Priority(rucksacks[1].Intersection()[0]))
	assert.Equal(t, 42, day3.Priority(rucksacks[2].Intersection()[0]))
	assert.Equal(t, 22, day3.Priority(rucksacks[3].Intersection()[0]))
	assert.Equal(t, 20, day3.Priority(rucksacks[4].Intersection()[0]))
	assert.Equal(t, 19, day3.Priority(rucksacks[5].Intersection()[0]))
	test := day3.AllRuckSacks(example)
	assert.Equal(t, 157, test.SumIntersectionPriorities())
}

func TestRucksackContent_Answer(t *testing.T) {
	input, err := readInput(t)
	require.NoError(t, err)

	allRuckSacks := day3.AllRuckSacks(input)
	t.Log("[part1] TestRucksackContent: ", allRuckSacks.SumIntersectionPriorities())
}

func TestGroupRucksackContent(t *testing.T) {

	test := day3.AllRuckSacks(example)
	assert.Equal(t, "r", test.Intersection(0))
	assert.Equal(t, 18, day3.Priority(test.Intersection(0)))
	assert.Equal(t, "Z", test.Intersection(1))
	assert.Equal(t, 52, day3.Priority(test.Intersection(1)))
	assert.Equal(t, 70, test.SumGroupIntersectionPriorities([]int{0, 1}))

}

func TestGroupRucksackContent_Answer(t *testing.T) {
	input, err := readInput(t)
	require.NoError(t, err)

	allRuckSacks := day3.AllRuckSacks(input)
	t.Log("[part2] TestGroupRucksackContent: ", allRuckSacks.SumAllGroupsIntersectionPriorities())
}

func readInput(t *testing.T) ([]string, error) {
	t.Helper()

	file, err := os.Open("./data/input.txt")
	require.NoError(t, err)
	defer file.Close()

	return parseInput(file)
}

func parseInput(f io.Reader) ([]string, error) {
	var (
		scanner = bufio.NewScanner(f)
		result  []string
	)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}
