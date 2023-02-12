package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestPart1(t *testing.T) {
	answer := p1(parse.String("input.test.txt"))
	if answer != 20899048083289 {
		t.Errorf("Expected 20899048083289 but was %d", answer)
	}
}
