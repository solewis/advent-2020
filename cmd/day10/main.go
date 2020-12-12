package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"sort"
)

func main() {
	jolts := parse.Ints("cmd/day10/input.txt", "\n")
	jolts = sortAndExpandJoltageList(jolts)
	fmt.Printf("Part 1: %d\n", findJoltageDistribution(jolts))
	fmt.Printf("Part 2: %d\n", findAdapterCombos(jolts))
}

func sortAndExpandJoltageList(jolts []int) []int {
	sort.Ints(jolts)
	jolts = append([]int{0}, jolts...)
	jolts = append(jolts, jolts[len(jolts)-1]+3)
	return jolts
}

func findJoltageDistribution(jolts []int) int {
	oneDiff, threeDiff := 0, 0
	for i := 0; i < len(jolts)-1; i++ {
		diff := jolts[i+1] - jolts[i]
		if diff == 1 {
			oneDiff++
		}
		if diff == 3 {
			threeDiff++
		}
	}
	return oneDiff * threeDiff
}

func findAdapterCombos(jolts []int) int {
	// break the adapters into smaller chains, where each chain represents a section that can branch in multiple ways
	// calculate the combos of each smaller chain and multiply them together to find the total number of combos
	var chains [][]int
	var currChain []int
	for i := 0; i < len(jolts)-1; i++ {
		currChain = append(currChain, jolts[i])
		if jolts[i+1]-jolts[i] == 3 {
			if len(currChain) > 2 {
				chains = append(chains, currChain)
			}
			currChain = nil
		}
	}
	product := 1
	for _, c := range chains {
		product *= findCombos(0, c)
	}

	return product
}

func findCombos(i int, jolts []int) int {
	// find total paths using dfs
	if i == len(jolts)-1 {
		return 1
	}
	sum := 0
	for j := i + 1; j < i+4; j++ {
		if j < len(jolts) && jolts[j]-jolts[i] <= 3 {
			sum += findCombos(j, jolts)
		}
	}
	return sum
}
