package main

import (
	"advent-2020/internal/parse"
	"reflect"
	"testing"
)

func TestParseProgram(t *testing.T) {
	program := parseProgram(parse.Lines("input.test.txt"))
	expected := []instruction{
		{"nop", 0}, {"acc", 1}, {"jmp", 4}, {"acc", 3}, {"jmp", -3},
		{"acc", -99}, {"acc", 1}, {"jmp", -4}, {"acc", 6},
	}
	if !reflect.DeepEqual(program, expected) {
		t.Errorf("Parse failed. Expected %v, got %v", expected, program)
	}
}

func TestRunProgram(t *testing.T) {
	program := parseProgram(parse.Lines("input.test.txt"))
	answer, _ := runProgram(program)
	if answer != 5 {
		t.Errorf("Expected 5 but was %d", answer)
	}
}

func TestAccAtTermination(t *testing.T) {
	program := parseProgram(parse.Lines("input.test.txt"))
	answer := accAtTermination(program)
	if answer != 8 {
		t.Errorf("Expected 8 but was %d", answer)
	}
}
