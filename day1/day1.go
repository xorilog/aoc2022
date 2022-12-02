package day1

import (
	"fmt"
	"strconv"
)

// SumCalories takes a slice of int and return the sum of all of them.
func SumCalories(l []int) int {
	var sum int
	for _, v := range l {
		sum = sum + v
	}
	return sum
}

// ProcessElfStats takes a list of string and treat a contiguous amount of valid int as an elf, it returns a slice of
// int with the sum of int per elf.
func ProcessElfStats(l []string) []int {
	var (
		leaderBoard   []int
		contiguousInt []int
	)

	for k, v := range l {
		if i, err := strconv.Atoi(v); err == nil {
			contiguousInt = append(contiguousInt, i)
		} else {
			leaderBoard = append(leaderBoard, SumCalories(contiguousInt))
			contiguousInt = nil
		}
		if k == len(l)-1 {
			fmt.Printf("Last round")
			leaderBoard = append(leaderBoard, SumCalories(contiguousInt))
		}
	}
	return leaderBoard
}

// GetPodium takes a slice of int l with an int s and return the top s from l as a slice of int
func GetPodium(l []int, s int) []int {
	var podium []int
	for i := 0; i < s; i++ {
		k, p := MaxFromSlice(l)
		podium = append(podium, p)
		l = RemoveIndex(l, k)
	}
	return podium
}

// GetPodiumSum takes a slice of string l and an int s and return the sum as int
func GetPodiumSum(l []string, s int) int {
	var sum int
	for _, v := range GetPodium(ProcessElfStats(l), s) {
		sum = sum + v
	}
	return sum
}

// RemoveIndex takes a slice s and an int i to return a slice without the index i
func RemoveIndex(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

// MaxFromSlice takes a list of int and return the index and highest value contained.
func MaxFromSlice(l []int) (int, int) {
	var highest, position int

	for k, v := range l {
		if v > highest {
			highest = v
			position = k
		}
	}

	return position, highest
}

// // HigestElf takes a slice of strings to extract the elf with the most calories
// func HigestElf(l []string) int {
// 	k, _ := MaxFromSlice(ProcessElfStats(l))
// 	return k + 1
// }

// HighestCalories takes a slice of strings to extract the maximum calories for an elf
func HighestCalories(l []string) int {
	_, v := MaxFromSlice(ProcessElfStats(l))
	return v
}
