package main

import (
	"strings"
	"testing"
)

func TestCalculateEqualPrecedence(t *testing.T) {
	t.Run("calculates line with + and * at the same precedence", func(t *testing.T) {
		line := strings.ReplaceAll("1 + 2 * 3 + 4 * 5 + 6", " ", "")
		answer := calculateEqualPrecedence(line)
		if answer != 71 {
			t.Errorf("Expected 71, but was %d", answer)
		}
	})

	t.Run("calculates parentheses first", func(t *testing.T) {
		line := strings.ReplaceAll("2 * 3 + (4 * 5)", " ", "")
		answer := calculateEqualPrecedence(line)
		if answer != 26 {
			t.Errorf("Expected 26, but was %d", answer)
		}

		line = strings.ReplaceAll("5 + (8 * 3 + 9 + 3 * 4 * 3)", " ", "")
		answer = calculateEqualPrecedence(line)
		if answer != 437 {
			t.Errorf("Expected 437, but was %d", answer)
		}
	})

	t.Run("handles nested parentheses", func(t *testing.T) {
		line := strings.ReplaceAll("1 + (2 * 3) + (4 * (5 + 6))", " ", "")
		answer := calculateEqualPrecedence(line)
		if answer != 51 {
			t.Errorf("Expected 51, but was %d", answer)
		}

		line = strings.ReplaceAll("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", " ", "")
		answer = calculateEqualPrecedence(line)
		if answer != 12240 {
			t.Errorf("Expected 12240, but was %d", answer)
		}
	})

	t.Run("handles parentheses at beginning", func(t *testing.T) {
		line := strings.ReplaceAll("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", " ", "")
		answer := calculateEqualPrecedence(line)
		if answer != 13632 {
			t.Errorf("Expected 13632, but was %d", answer)
		}
	})
}
func TestCalculateAdditionFirst(t *testing.T) {
	t.Run("calculates line with + before *", func(t *testing.T) {
		line := strings.ReplaceAll("1 + 2 * 3 + 4 * 5 + 6", " ", "")
		answer := calculateAdditionFirst(line)
		if answer != 231 {
			t.Errorf("Expected 231, but was %d", answer)
		}
	})

	t.Run("handles parentheses first", func(t *testing.T) {
		line := strings.ReplaceAll("1 + (2 * 3) + (4 * (5 + 6))", " ", "")
		answer := calculateAdditionFirst(line)
		if answer != 51 {
			t.Errorf("Expected 51, but was %d", answer)
		}

		line = strings.ReplaceAll("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", " ", "")
		answer = calculateAdditionFirst(line)
		if answer != 669060 {
			t.Errorf("Expected 669060, but was %d", answer)
		}
	})

	t.Run("solves the leftmost addition first", func(t *testing.T) {
		//tests an edge case where solving the rightmost addition first will fail
		//caught a bug with regex
		line := strings.ReplaceAll("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", " ", "")
		answer := calculateAdditionFirst(line)
		if answer != 23340 {
			t.Errorf("Expected 23340, but was %d", answer)
		}
	})
}