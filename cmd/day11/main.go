package main

import (
	"advent-2020/internal/parse"
	"fmt"
)

func main() {
	lines := parse.Lines("cmd/day11/input.txt")
	g := parseGrid(lines)
	fmt.Printf("Part 1: %d\n", countOccupiedSeats(g, 4, countOccupiedAdjacent))
	fmt.Printf("Part 2: %d\n", countOccupiedSeats(g, 5, countOccupiedLineOfSight))
}

type grid [][]rune

func (g grid) inGrid(row, col int) bool {
	return row >= 0 && row < len(g) && col >= 0 && col < len(g[0])
}

func (g grid) isVal(val rune, row, col int) bool {
	return g.inGrid(row, col) && g[row][col] == val
}

func (g grid) countVal(val rune) int {
	count := 0
	for row := range g {
		for col := range g[row] {
			if g.isVal(val, row, col) {
				count++
			}
		}
	}
	return count
}

func (g grid) String() string {
	s := ""
	for row := range g {
		l := ""
		for col := range g[row] {
			l += string(g[row][col])
		}
		s += l + "\n"
	}
	return s
}

func parseGrid(lines []string) grid {
	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}
	return grid
}

func countOccupiedSeats(grid grid, occupiedSeatAmt int, ruleFunc func(grid, int, int) int) int {
	for changed := true; changed; {
		grid, changed = simulateRound(grid, occupiedSeatAmt, ruleFunc)
	}
	return grid.countVal('#')
}

func simulateRound(prevGrid grid, occupiedSeatAmt int, ruleFunc func(grid, int, int) int) (grid, bool) {
	nextGrid := make([][]rune, len(prevGrid))
	for i := range nextGrid {
		nextGrid[i] = make([]rune, len(prevGrid[0]))
	}
	changed := false
	for row := range prevGrid {
		for col := range prevGrid[row] {
			val := prevGrid[row][col]
			adj := ruleFunc(prevGrid, row, col)

			switch {
			case val == 'L' && adj == 0:
				nextGrid[row][col] = '#'
				changed = true
			case val == '#' && adj >= occupiedSeatAmt:
				nextGrid[row][col] = 'L'
				changed = true
			default:
				nextGrid[row][col] = val
			}
		}
	}
	return nextGrid, changed
}

func countOccupiedAdjacent(grid grid, row, col int) int {
	count := 0
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if r == row && c == col {
				continue
			}
			if grid.isVal('#', r, c) {
				count++
			}
		}
	}
	return count
}

func countOccupiedLineOfSight(grid grid, row, col int) int {
	lineDirs := []struct{ rowDir, colDir int }{
		{0, -1},  //left
		{0, 1},   //right
		{-1, 0},  //up
		{1, 0},   //down
		{-1, -1}, //up left
		{-1, 1},  //up right
		{1, -1},  //down left
		{1, 1},   //down right
	}

	count := 0
	for _, d := range lineDirs {
		if checkLineOccupied(grid, row, col, d.rowDir, d.colDir) {
			count++
		}
	}
	return count
}

func checkLineOccupied(grid grid, row, col, rowDir, colDir int) bool {
	row += rowDir
	col += colDir
	for grid.inGrid(row, col) {
		if grid.isVal('L', row, col) {
			return false
		}
		if grid.isVal('#', row, col) {
			return true
		}
		row += rowDir
		col += colDir
	}
	return false
}
