package day2_test

import (
	"bufio"
	"github.com/xorilog/aoc2022/day2"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var example = []string{
	"A Y",
	"B X",
	"C Z",
}

//
// func TestHighestElf(t *testing.T) {
// 	elfNumber := day2.HighestElf(example)
// 	assert.Equal(t, 4, elfNumber)
// }

// func TestHighestElf_Answer_Part1(t *testing.T) {
// 	input, err := readInput(t)
// 	require.NoError(t, err)
//
// 	t.Log("[part1] HighestElf: ", day2.HighestElf(input))
// }

func TestRockPaperScissors(t *testing.T) {
	points := day2.SumOfGames(example)
	assert.Equal(t, 15, points)
}

func TestRockPaperScissors_Answer(t *testing.T) {
	input, err := readInput(t)
	require.NoError(t, err)

	t.Log("[part1] RockPaperScissors: ", day2.SumOfGames(input))
}

func TestRockPaperScissorsRoundBasedStrategy(t *testing.T) {
	points := day2.SumOfGamesRoundStrategy(example)
	assert.Equal(t, 12, points)
}

func TestRockPaperScissorsRoundBasedStrategy_Answer(t *testing.T) {
	input, err := readInput(t)
	require.NoError(t, err)

	t.Log("[part2] RockPaperScissorsRoundBasedStrategy: ", day2.SumOfGamesRoundStrategy(input))
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
