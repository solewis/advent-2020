package main

import (
	"advent-2020/internal/parse"
	"fmt"
)

func main() {
	lines := parse.Lines("cmd/day17/input.txt")
	fmt.Printf("Part 1: %d\n", p1(lines))
	fmt.Printf("Part 2: %d\n", p2(lines))
}

func p1(lines []string) int {
	grid := parseBootGrid(lines)
	for i := 0; i < 6; i++ {
		grid = grid.runCycle()
	}
	return grid.activeCount()
}

type point struct {
	x, y, z int
}

//returns all 26 neighbors of the given point
func (p point) neighbors() []point {
	var neighbors []point
	for x := p.x - 1; x <= p.x+1; x++ {
		for y := p.y - 1; y <= p.y+1; y++ {
			for z := p.z - 1; z <= p.z+1; z++ {
				neighbor := point{x, y, z}
				if neighbor != p {
					neighbors = append(neighbors, neighbor)
				}
			}
		}
	}
	return neighbors
}

type bootGrid struct {
	grid                               map[point]bool
	minX, maxX, minY, maxY, minZ, maxZ int
}

//runs a single cycle of the boot process on the grid, returning the next iteration of the grid
func (g bootGrid) runCycle() bootGrid {
	nextGrid := bootGrid{
		grid: map[point]bool{},
		minX: g.minX - 1,
		maxX: g.maxX + 1,
		minY: g.minY - 1,
		maxY: g.maxY + 1,
		minZ: g.minZ - 1,
		maxZ: g.maxZ + 1,
	}

	for x := nextGrid.minX; x <= nextGrid.maxX; x++ {
		for y := nextGrid.minY; y <= nextGrid.maxY; y++ {
			for z := nextGrid.minZ; z <= nextGrid.maxZ; z++ {
				current := point{x, y, z}
				currentActive := g.grid[current]
				activeNeighbors := g.activeNeighbors(current)
				if currentActive && (activeNeighbors == 2 || activeNeighbors == 3) {
					nextGrid.grid[current] = true
				} else if !currentActive && activeNeighbors == 3 {
					nextGrid.grid[current] = true
				} else {
					nextGrid.grid[current] = false
				}
			}
		}
	}

	return nextGrid
}

//returns the number of active points in the grid
func (g bootGrid) activeCount() int {
	activeCount := 0
	for _, active := range g.grid {
		if active {
			activeCount++
		}
	}
	return activeCount
}

//returns the number of active neighbors of a given point on the grid
func (g bootGrid) activeNeighbors(p point) int {
	count := 0
	for _, n := range p.neighbors() {
		if g.grid[n] {
			count++
		}
	}
	return count
}

func parseBootGrid(lines []string) bootGrid {
	bg := bootGrid{
		grid: map[point]bool{},
		minX: 0,
		maxX: len(lines[0]),
		minY: 0,
		maxY: len(lines),
		minZ: 0,
		maxZ: 0,
	}
	for y, line := range lines {
		for x, val := range line {
			isActive := false
			if val == '#' {
				isActive = true
			}
			bg.grid[point{x, y, 0}] = isActive
		}
	}
	return bg
}
