package main

import "testing"

func TestDistanceToSafety(t *testing.T) {
	answer := distanceToSafety([]string{"F10", "N3", "F7", "R90", "F11"})
	if answer != 25 {
		t.Errorf("Expected 25 but was %d", answer)
	}
}

func TestActualDistanceToSafety(t *testing.T) {
	answer := actualDistanceToSafety([]string{"F10", "N3", "F7", "R90", "F11"})
	if answer != 286 {
		t.Errorf("Expected 286 but was %d", answer)
	}
}
