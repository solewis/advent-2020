package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"testing"
)

func TestFindJoltageDistribution(t *testing.T) {
	jolts := parse.Ints("input.test.txt", "\n")
	dist := findJoltageDistribution(jolts)
	if dist != 220 {
		t.Errorf("Expected 220 but was %d", dist)
	}
}

func TestFindAdapterCombos(t *testing.T) {
	t.Run("brute", func(t *testing.T) {
		t.Run("sample 1", func(t *testing.T) {
			jolts := parse.Ints("input.test.txt", "\n")
			jolts = sortAndExpandJoltageList(jolts)
			combos := findCombos(0, jolts)
			if combos != 19208 {
				t.Errorf("Expected 19208 but was %d", combos)
			}
		})

		t.Run("sample 2", func(t *testing.T) {
			jolts := parse.Ints("input.test2.txt", "\n")
			jolts = sortAndExpandJoltageList(jolts)
			combos := findCombos(0, jolts)
			if combos != 8 {
				t.Errorf("Expected 8 but was %d", combos)
			}
		})
	})

	t.Run("fast", func(t *testing.T) {
		t.Run("sample 1", func(t *testing.T) {
			jolts := parse.Ints("input.test.txt", "\n")
			jolts = sortAndExpandJoltageList(jolts)
			fmt.Println(jolts)
			combos := findAdapterCombos(jolts)
			if combos != 19208 {
				t.Errorf("Expected 19208 but was %d", combos)
			}
		})

		t.Run("sample 2", func(t *testing.T) {
			jolts := parse.Ints("input.test2.txt", "\n")
			jolts = sortAndExpandJoltageList(jolts)
			fmt.Println(jolts)
			combos := findAdapterCombos(jolts)
			if combos != 8 {
				t.Errorf("Expected 8 but was %d", combos)
			}
		})
	})

	t.Run("bottom up", func(t *testing.T) {
		t.Run("sample 1", func(t *testing.T) {
			jolts := parse.Ints("input.test.txt", "\n")
			jolts = sortAndExpandJoltageList(jolts)
			fmt.Println(jolts)
			combos := findAdapterCombosBottomUp(jolts)
			if combos != 19208 {
				t.Errorf("Expected 19208 but was %d", combos)
			}
		})

		t.Run("sample 2", func(t *testing.T) {
			jolts := parse.Ints("input.test2.txt", "\n")
			jolts = sortAndExpandJoltageList(jolts)
			fmt.Println(jolts)
			combos := findAdapterCombosBottomUp(jolts)
			if combos != 8 {
				t.Errorf("Expected 8 but was %d", combos)
			}
		})
	})
}
