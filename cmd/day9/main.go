package main

import (
	"advent-2020/internal/parse"
	"fmt"
)

func main() {
	data := parse.Ints("cmd/day9/input.txt", "\n")
	firstInvalid := findFirstInvalid(data, 25)
	fmt.Printf("Part 1: %d\n", firstInvalid)
	fmt.Printf("Part 2: %d\n", findEncryptionWeakness(firstInvalid, data))
}

func findFirstInvalid(data []int, preambleLength int) int {
	for i := 0; true; i++ {
		if !isValid(data[i+preambleLength], data[i:i+preambleLength]) {
			return data[i+preambleLength]
		}
	}
	return -1
}

func findEncryptionWeakness(num int, data []int) int {
	for i := range data {
		sum := data[i]
		for j := i + 1; j < len(data); j++ {
			sum += data[j]
			if sum == num {
				return sumMinMaxInRange(data[i:j + 1])
			}
		}
	}
	return -1
}

func sumMinMaxInRange(r []int) int {
	min, max := r[0], r[0]
	for _, i := range r {
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}
	return min + max
}

func isValid(num int, preamble []int) bool {
	for i := range preamble {
		for j := i + 1; j < len(preamble); j++ {
			if preamble[i] != preamble[j] && preamble[i]+preamble[j] == num {
				return true
			}
		}
	}
	return false
}
