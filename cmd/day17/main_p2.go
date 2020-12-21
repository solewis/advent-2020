package main

func p2(lines []string) int {
	grid := parseBootGrid4d(lines)
	for i := 0; i < 6; i++ {
		grid = grid.runCycle()
	}
	return grid.activeCount()
}

type point4d struct {
	x, y, z, w int
}

//returns all 80 neighbors of the given point
func (p point4d) neighbors() []point4d {
	var neighbors []point4d
	for x := p.x - 1; x <= p.x+1; x++ {
		for y := p.y - 1; y <= p.y+1; y++ {
			for z := p.z - 1; z <= p.z+1; z++ {
				for w := p.w - 1; w <= p.w+1; w++ {
					neighbor := point4d{x, y, z, w}
					if neighbor != p {
						neighbors = append(neighbors, neighbor)
					}
				}
			}
		}
	}
	return neighbors
}

type bootGrid4d struct {
	grid                                           map[point4d]bool
	minX, maxX, minY, maxY, minZ, maxZ, minW, maxW int
}

//runs a single cycle of the boot process on the grid, returning the next iteration of the grid
func (g bootGrid4d) runCycle() bootGrid4d {
	nextGrid := bootGrid4d{
		grid: map[point4d]bool{},
		minX: g.minX - 1,
		maxX: g.maxX + 1,
		minY: g.minY - 1,
		maxY: g.maxY + 1,
		minZ: g.minZ - 1,
		maxZ: g.maxZ + 1,
		minW: g.minW - 1,
		maxW: g.maxW + 1,
	}

	for x := nextGrid.minX; x <= nextGrid.maxX; x++ {
		for y := nextGrid.minY; y <= nextGrid.maxY; y++ {
			for z := nextGrid.minZ; z <= nextGrid.maxZ; z++ {
				for w := nextGrid.minW; w <= nextGrid.maxW; w++ {
					current := point4d{x, y, z, w}
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
	}

	return nextGrid
}

//returns the number of active points in the grid
func (g bootGrid4d) activeCount() int {
	activeCount := 0
	for _, active := range g.grid {
		if active {
			activeCount++
		}
	}
	return activeCount
}

//returns the number of active neighbors of a given point on the grid
func (g bootGrid4d) activeNeighbors(p point4d) int {
	count := 0
	for _, n := range p.neighbors() {
		if g.grid[n] {
			count++
		}
	}
	return count
}

func parseBootGrid4d(lines []string) bootGrid4d {
	bg := bootGrid4d{
		grid: map[point4d]bool{},
		minX: 0,
		maxX: len(lines[0]),
		minY: 0,
		maxY: len(lines),
		minZ: 0,
		maxZ: 0,
		minW: 0,
		maxW: 0,
	}
	for y, line := range lines {
		for x, val := range line {
			isActive := false
			if val == '#' {
				isActive = true
			}
			bg.grid[point4d{x, y, 0, 0}] = isActive
		}
	}
	return bg
}
