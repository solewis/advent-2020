package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestDay17(t *testing.T) {
	t.Run("Test Grid", func(t *testing.T) {
		lines := parse.Lines("input.test.txt")
		grid := parseBootGrid(lines)
		if grid.activeCount() != 5 {
			t.Errorf("Expected 5 active after parsing, but was %d", grid.activeCount())
		}
		grid = grid.runCycle()
		if grid.activeCount() != 11 {
			t.Errorf("Expected 11 active 1 cycle, but was %d", grid.activeCount())
		}
		grid = grid.runCycle()
		if grid.activeCount() != 21 {
			t.Errorf("Expected 21 active after 2 cycles, but was %d", grid.activeCount())
		}
		grid = grid.runCycle().runCycle().runCycle().runCycle()
		if grid.activeCount() != 112 {
			t.Errorf("Expected 112 active after 6 cycles, but was %d", grid.activeCount())
		}
	})

	t.Run("Test 4d Grid", func(t *testing.T) {
		lines := parse.Lines("input.test.txt")
		grid := parseBootGrid4d(lines)
		if grid.activeCount() != 5 {
			t.Errorf("Expected 5 active after parsing, but was %d", grid.activeCount())
		}
		grid = grid.runCycle()
		if grid.activeCount() != 29 {
			t.Errorf("Expected 29 active 1 cycle, but was %d", grid.activeCount())
		}
		grid = grid.runCycle().runCycle().runCycle().runCycle().runCycle()
		if grid.activeCount() != 848 {
			t.Errorf("Expected 848 active after 6 cycles, but was %d", grid.activeCount())
		}
	})
}
