package main

import (
	"advent-2020/internal/parse"
	"fmt"
)

func main() {
	instructions := parse.Lines("cmd/day12/input.txt")
	fmt.Printf("Part 1: %d\n", distanceToSafety(instructions))
	fmt.Printf("Part 2: %d\n", actualDistanceToSafety(instructions))
}

const (
	North = iota
	East
	South
	West
)

func runeVal(direction int) rune {
	return [...]rune{'N', 'E', 'S', 'W'}[direction]
}

func distanceToSafety(instructions []string) int {
	x, y, facing := 0, 0, East
	for _, ins := range instructions {
		action := rune(ins[0])
		value := parse.Int(ins[1:])
		if action == 'F' {
			action = runeVal(facing)
		}
		switch action {
		case 'N':
			y += value
		case 'S':
			y -= value
		case 'E':
			x += value
		case 'W':
			x -= value
		case 'L':
			facing = (facing + 4 - value/90) % 4
		case 'R':
			facing = (facing + value/90) % 4
		}
	}

	return abs(x) + abs(y)
}

func actualDistanceToSafety(instructions []string) int {
	sx, sy, wx, wy := 0, 0, 10, 1
	for _, ins := range instructions {
		action := rune(ins[0])
		value := parse.Int(ins[1:])
		switch action {
		case 'N':
			wy += value
		case 'S':
			wy -= value
		case 'E':
			wx += value
		case 'W':
			wx -= value
		case 'L':
			for i := 0; i < value/90; i++ {
				tmp := wx
				wx = -wy
				wy = tmp
			}
		case 'R':
			for i := 0; i < value/90; i++ {
				tmp := wx
				wx = wy
				wy = -tmp
			}
		case 'F':
			sx += wx * value
			sy += wy * value
		}

	}

	return abs(sx) + abs(sy)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
