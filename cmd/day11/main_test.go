package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestCountOccupiedWithAdjacencyRule(t *testing.T) {
	lines := parse.Lines("input.test.txt")
	grid := parseGrid(lines)
	answer := countOccupiedSeats(grid, 4, countOccupiedAdjacent)
	if answer != 37 {
		t.Errorf("Expected 37 occupied seats, but was %d", answer)
	}
}

func TestCountOccupiedWithLineOfSightRule(t *testing.T) {
	t.Run("line of sight rule", func(t *testing.T) {
		t.Run("sees along line", func(t *testing.T) {
			grid := [][]rune{
				{'.', '#', '.'},
				{'.', '.', '.'},
				{'.', '#', '.'},
			}
			answer := countOccupiedLineOfSight(grid, 2, 1)
			if answer != 1 {
				t.Errorf("Expected 1 occupied seat, but was %d", answer)
			}
		})

		t.Run("cannot see past empty seat", func(t *testing.T) {
			grid := [][]rune{
				{'.', '#', '.'},
				{'.', 'L', '.'},
				{'.', '#', '.'},
			}
			answer := countOccupiedLineOfSight(grid, 2, 1)
			if answer != 0 {
				t.Errorf("Expected 0 occupied seat, but was %d", answer)
			}
		})
	})

	t.Run("full count", func(t *testing.T) {
		lines := parse.Lines("input.test.txt")
		grid := parseGrid(lines)
		answer := countOccupiedSeats(grid, 5, countOccupiedLineOfSight)
		if answer != 26 {
			t.Errorf("Expected 26 occupied seats, but was %d", answer)
		}
	})
}
