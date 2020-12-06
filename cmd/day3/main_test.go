package main

import "testing"

func TestTreeOnMap(t *testing.T) {
	m := forest{data: [][]rune{{'.', '#', '.', '#'}, {'#', '#', '.', '#'}, {'.', '.', '.', '.'}, {'.', '#', '#', '#'}}}
	t.Run("finds tree on map", func(t *testing.T) {
		if m.isTree(point{0, 0}) {
			t.Error("Expected point 0, 0 to not be a tree")
		}
		if !m.isTree(point{0, 1}) {
			t.Error("Expected point 0, 1 to be a tree")
		}
	})
	t.Run("map repeats on the x axis", func(t *testing.T) {
		if m.isTree(point{4, 0}) {
			t.Error("Expected point 5, 0 to not be a tree")
		}
		if !m.isTree(point{4, 1}) {
			t.Error("Expected point 5, 1 to be a tree")
		}
	})
}
