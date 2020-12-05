package main

import (
	"advent-2020/internal/parse"
	"fmt"
)

func main() {
	entries := parse.Ints("cmd/day1/input.txt", "\n")
	i, j := twoEntriesThatSumTo2020(entries)
	fmt.Printf("Part 1: %d\n", entries[i]*entries[j])

	i, j, k := threeEntriesThatSumTo2020(entries)
	fmt.Printf("Part 2: %d\n", entries[i]*entries[j]*entries[k])
}

func twoEntriesThatSumTo2020(entries []int) (int, int) {
	for i := range entries {
		for j := i + 1; j < len(entries); j++ {
			if entries[i]+entries[j] == 2020 {
				return i, j
			}
		}
	}
	return -1, -1
}

func threeEntriesThatSumTo2020(entries []int) (int, int, int) {
	for i := range entries {
		for j := i + 1; j < len(entries); j++ {
			for k := j + 1; k < len(entries); k++ {
				if entries[i]+entries[j]+entries[k] == 2020 {
					return i, j, k
				}
			}
		}
	}
	return -1, -1, -1
}
