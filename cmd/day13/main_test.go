package main

import "testing"

func TestPart1(t *testing.T) {
	timestamp := 939
	ids := []bus{{0, 7}, {1, 13}, {2, 59}, {3, 31}, {4, 19}}
	answer := part1(timestamp, ids)
	if answer != 295 {
		t.Errorf("Expected 295 but was %d", answer)
	}
}

func TestPart2(t *testing.T) {
	t.Run("sample 1", func(t *testing.T) {
		ids := []bus{{0, 7}, {1, 13}, {4, 59}, {6, 31}, {7, 19}}
		answer := part2(ids)
		if answer != 1068781 {
			t.Errorf("Expected 1068781 but was %d", answer)
		}
	})
	t.Run("sample 6", func(t *testing.T) {
		ids := []bus{{0, 1789}, {1, 37}, {2, 47}, {3, 1889}}
		answer := part2(ids)
		if answer != 1202161486 {
			t.Errorf("Expected 1202161486 but was %d", answer)
		}
	})
}
