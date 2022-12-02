package day1_test

import (
	"bufio"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xorilog/aoc2022/day1"
)

var example = []string{
	"1000",
	"2000",
	"3000",
	"",
	"4000",
	"",
	"5000",
	"6000",
	"",
	"7000",
	"8000",
	"9000",
	"",
	"10000",
}

//
// func TestHighestElf(t *testing.T) {
// 	elfNumber := day1.HighestElf(example)
// 	assert.Equal(t, 4, elfNumber)
// }

// func TestHighestElf_Answer_Part1(t *testing.T) {
// 	input, err := readInput(t)
// 	require.NoError(t, err)
//
// 	t.Log("[part1] HighestElf: ", day1.HighestElf(input))
// }

func TestHighestCalories(t *testing.T) {
	calories := day1.HighestCalories(example)
	assert.Equal(t, 24000, calories)
}

func TestHighestCalories_Answer(t *testing.T) {
	input, err := readInput(t)
	require.NoError(t, err)

	t.Log("[part1] HighestCalories: ", day1.HighestCalories(input))
}

func TestTop3HighestCalories(t *testing.T) {
	calories := day1.GetPodiumSum(example, 3)
	assert.Equal(t, 45000, calories)
}

func TestTop3HighestCalories_Answer(t *testing.T) {
	input, err := readInput(t)
	require.NoError(t, err)

	t.Log("[part2] Top3HighestCalories: ", day1.GetPodiumSum(input, 3))
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
