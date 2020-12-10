package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"regexp"
	"strings"
)

var (
	valueRegex = regexp.MustCompile(`^(\d+) (.+) bags?\.?$`)
)

func main() {
	lines := parse.Lines("cmd/day7/input.txt")
	rules := parseRules(lines)
	fmt.Printf("Part 1: %d\n", countBagsThatCanHold("shiny gold", rules, map[string]bool{}))
	fmt.Printf("Part 2: %d\n", countBagsRequiredInside("shiny gold", rules))
}

type contents map[string]int
type rules map[string]contents

func countBagsThatCanHold(subject string, rules rules, acc map[string]bool) int {
	for k, v := range rules {
		if _, ok := v[subject]; ok {
			acc[k] = true
			countBagsThatCanHold(k, rules, acc)
		}
	}
	return len(acc)
}

func countBagsRequiredInside(subject string, rules rules) int {
	count := 0
	for k, v := range rules[subject] {
		count += countBagsRequiredInside(k, rules) * v + v
	}
	return count
}

func parseRules(lines []string) rules {
	ruleMap := rules{}
	for _, line := range lines {
		parts := strings.Split(line, " bags contain ")
		valueParts := strings.Split(parts[1], ", ")
		value := contents{}
		for _, vp := range valueParts {
			if vp != "no other bags." {
				matches := valueRegex.FindStringSubmatch(vp)
				value[matches[2]] = parse.Int(matches[1])
			}
		}
		ruleMap[parts[0]] = value
	}
	return ruleMap
}
