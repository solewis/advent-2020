package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestSumOfFinalMemoryV1(t *testing.T) {
	instructions := parse.Lines("input.test.txt")
	answer := sumOfFinalMemory(instructions, runProgram)
	if answer != 165 {
		t.Errorf("Expected 165, but was %d", answer)
	}
}

func TestSumOfFinalMemoryV2(t *testing.T) {
	instructions := parse.Lines("input.test2.txt")
	answer := sumOfFinalMemory(instructions, runProgramV2)
	if answer != 208 {
		t.Errorf("Expected 208, but was %d", answer)
	}
}
