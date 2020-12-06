package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"sort"
)

func main() {
	lines := parse.Lines("cmd/day5/input.txt")
	fmt.Printf("Part 1: %d\n", findMaxSeatId(lines))
	fmt.Printf("Part 2: %d\n", findMySeat(lines))
}

func findMySeat(lines []string) int {
	seats := parseSeats(lines)
	sort.Ints(seats)
	for i := range seats {
		if seats[i + 1] > seats[i] + 1 {
			return seats[i] + 1
		}
	}
	return -1
}

func findMaxSeatId(lines []string) int {
	max := 0
	seats := parseSeats(lines)
	for _, seat := range seats {
		if seat > max {
			max = seat
		}
	}
	return max
}

func parseSeats(lines []string) []int {
	var seats []int
	for _, line := range lines {
		rowData := line[:7]
		columnData := line[7:]
		row := traverseInstruction(rowData, 0, 127)
		column := traverseInstruction(columnData, 0, 7)
		seats = append(seats, row * 8 + column)
	}
	return seats
}


func traverseInstruction(data string, min, max int) int {
	if data == "" {
		return min
	}
	if data[0] == 'F' || data[0] == 'L' {
		return traverseInstruction(data[1:], min, ((max - min) / 2) + min)
	}
	return traverseInstruction(data[1:], (((max - min) / 2) + 1) + min, max)
}