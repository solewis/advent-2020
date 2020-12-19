package main

import (
	"fmt"
	"testing"
	"time"
)

func TestRunGame(t *testing.T) {
	runTest := func(t *testing.T, seed []int, length, expected int) {
		answer := runGame(seed, length)
		if answer != expected {
			t.Errorf("Expected %d, but was %d", expected, answer)
		}
	}
	t.Run("game 1", func(t *testing.T) {
		runTest(t, []int{0, 3, 6}, 2020, 436)
		now := time.Now()
		runTest(t, []int{0, 3, 6}, 30000000, 175594)
		fmt.Println(time.Since(now))
	})
	t.Run("game 2", func(t *testing.T) {
		runTest(t, []int{1, 3, 2}, 2020, 1)
	})
	t.Run("game 3", func(t *testing.T) {
		runTest(t, []int{2, 1, 3}, 2020, 10)
	})
	t.Run("game 4", func(t *testing.T) {
		runTest(t, []int{1, 2, 3}, 2020, 27)
	})
	t.Run("game 5", func(t *testing.T) {
		runTest(t, []int{2, 3, 1}, 2020, 78)
	})
	t.Run("game 6", func(t *testing.T) {
		runTest(t, []int{3, 2, 1}, 2020, 438)
	})
	t.Run("game 7", func(t *testing.T) {
		runTest(t, []int{3, 1, 2}, 2020, 1836)
	})
}
