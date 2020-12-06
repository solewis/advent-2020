package main

import (
	"advent-2020/internal/parse"
	"fmt"
)

func main() {
	lines := parse.Lines("cmd/day3/input.txt")
	f := parseForest(lines)
	fmt.Printf("Part 1: %d\n", f.treesOnPath(3, 1))

	t1 := f.treesOnPath(1, 1)
	t2 := f.treesOnPath(3, 1)
	t3 := f.treesOnPath(5, 1)
	t4 := f.treesOnPath(7, 1)
	t5 := f.treesOnPath(1, 2)

	fmt.Printf("Part 2: %d\n", t1*t2*t3*t4*t5)
}

func parseForest(lines []string) forest {
	data := make([][]rune, len(lines))
	for i, line := range lines {
		currentRow := make([]rune, len(line))
		for j, r := range line {
			currentRow[j] = r
		}
		data[i] = currentRow
	}
	return forest{data}
}

type forest struct {
	data [][]rune
}

func (f forest) isTree(x, y int) bool {
	return f.data[y][x%len(f.data[0])] == '#'
}

func (f forest) treesOnPath(xMove, yMove int) int {
	currentX, currentY, trees := 0, 0, 0
	for currentY < len(f.data) {
		if f.isTree(currentX, currentY) {
			trees++
		}
		currentX += xMove
		currentY += yMove
	}
	return trees
}
