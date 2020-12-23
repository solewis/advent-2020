package main

import (
	"advent-2020/internal/parse"
	"strings"
	"testing"
)

func TestRuleMap(t *testing.T) {
	t.Run("sample 1", func(t *testing.T) {
		data := parse.String("input.test.txt")
		parts := strings.Split(data, "\n\n")
		ruleMap := parseRules(strings.Split(parts[0], "\n"))
		answer := ruleMap.countMessagesMatchingRule(0, strings.Split(parts[1], "\n"))
		if answer != 2 {
			t.Errorf("Expected 2 messages to match rule 0, but was %d", answer)
		}
	})

	t.Run("sample 2", func(t *testing.T) {
		data := parse.String("input.test2.txt")
		parts := strings.Split(data, "\n\n")
		ruleMap := parseRules(strings.Split(parts[0], "\n"))
		messages := strings.Split(parts[1], "\n")
		answer := ruleMap.countMessagesMatchingRule(0, messages)
		if answer != 3 {
			t.Errorf("Expected 3 messages to match rule 0, but was %d", answer)
		}
	})

	t.Run("test part 2", func(t *testing.T) {
		data := parse.String("input.test2.txt")
		parts := strings.Split(data, "\n\n")
		ruleMap := parseRules(strings.Split(parts[0], "\n"))
		messages := strings.Split(parts[1], "\n")
		answer := p2(ruleMap, messages)
		if answer != 12 {
			t.Errorf("Expected 12 messages to match rule 0, but was %d", answer)
		}
	})
}
