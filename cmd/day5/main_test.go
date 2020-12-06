package main

import "testing"

func TestTraverse(t *testing.T) {
	row := traverseInstruction("FBFBBFF", 0, 127)
	if row != 44 {
		t.Error("Expected row 44")
	}
	column := traverseInstruction("RLR", 0, 7)
	if column != 5 {
		t.Error("Expected column 5")
	}
}
