package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"testing"
)

func TestParseRules(t *testing.T) {
	rules := parseRules(parse.Lines("input.test.txt"))
	expected := map[string]map[string]int{
		"light red": {"bright white": 1, "muted yellow": 2},
		"dark orange": {"bright white": 3, "muted yellow": 4},
		"bright white": {"shiny gold": 1},
		"muted yellow": {"shiny gold": 2, "faded blue": 9},
		"shiny gold": {"dark olive": 1, "vibrant plum": 2},
		"dark olive": {"faded blue": 3, "dotted black": 4},
		"vibrant plum": {"faded blue": 5, "dotted black": 6},
		"faded blue": nil,
		"dotted black": nil,
	}
	if fmt.Sprint(rules) != fmt.Sprint(expected) {
		t.Errorf("parse failed: %s not equal to %s", fmt.Sprint(rules), fmt.Sprint(expected))
	}
}

func TestCountBagsThatCanHold(t *testing.T) {
	rules := parseRules(parse.Lines("input.test.txt"))
	answer := countBagsThatCanHold("shiny gold", rules, map[string]bool{})
	if answer != 4 {
		t.Errorf("Expected 4 bags to be able to carry shiny gold, but was %d", answer)
	}
}

func TestBagsRequiredInside(t *testing.T) {
	rules := parseRules(parse.Lines("input.test.txt"))
	answer := countBagsRequiredInside("shiny gold", rules)
	if answer != 32 {
		t.Errorf("Expected 32 bags to be able to carry shiny gold, but was %d", answer)
	}

	rules = parseRules(parse.Lines("input.test2.txt"))
	answer = countBagsRequiredInside("shiny gold", rules)
	if answer != 126 {
		t.Errorf("Expected 126 bags to be able to carry shiny gold, but was %d", answer)
	}
}
