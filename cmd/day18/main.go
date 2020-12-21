package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := parse.Lines("cmd/day18/input.txt")
	fmt.Printf("Part 1: %d\n", sumOfCalculatedLines(lines, calculateEqualPrecedence))
	fmt.Printf("Part 2: %d\n", sumOfCalculatedLines(lines, calculateAdditionFirst))
}

func sumOfCalculatedLines(lines []string, solveFunc func(string) int) int {
	sum := 0
	for _, line := range lines {
		line = strings.ReplaceAll(line, " ", "")
		sum += solveFunc(line)
	}
	return sum
}

var (
	//matches the first + or * operation
	firstOpRegex  = regexp.MustCompile(`^(\d+)([*+])(\d+)(.*)$`)
	//matches the first + operation
	firstAddRegex = regexp.MustCompile(`^(.*?)(\d+)\+(\d+)(.*)$`)
	//matches the first set of parentheses that do not have any parentheses inside it
	parensRegex   = regexp.MustCompile(`^(.*?)\(([0-9+*]+)\)(.*)$`)
)

//order of precedence: 1 - parenthesis, 2 - addition and multiplication
func calculateEqualPrecedence(line string) int {
	if val, err := strconv.Atoi(line); err == nil {
		return val
	}
	line = expandParenthesis(line, calculateEqualPrecedence)

	matches := firstOpRegex.FindStringSubmatch(line)
	a := parse.Int(matches[1])
	b := parse.Int(matches[3])
	if matches[2] == "*" {
		return calculateEqualPrecedence(strconv.Itoa(a*b) + matches[4])
	}
	return calculateEqualPrecedence(strconv.Itoa(a+b) + matches[4])
}

//order of precedence: 1 - parenthesis, 2 - addition, 3 - multiplication
func calculateAdditionFirst(line string) int {
	if val, err := strconv.Atoi(line); err == nil {
		return val
	}
	line = expandParenthesis(line, calculateAdditionFirst)

	if matches := firstAddRegex.FindStringSubmatch(line); matches != nil {
		a := parse.Int(matches[2])
		b := parse.Int(matches[3])
		return calculateAdditionFirst(matches[1] + strconv.Itoa(a+b) + matches[4])
	}
	return calculateEqualPrecedence(line)
}

func expandParenthesis(line string, solveFunc func(string) int) string {
	if matches := parensRegex.FindStringSubmatch(line); matches != nil {
		return expandParenthesis(matches[1]+strconv.Itoa(solveFunc(matches[2]))+matches[3], solveFunc)
	}
	return line
}
