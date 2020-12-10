package main

import (
	"advent-2020/internal/parse"
	"fmt"
)

func main() {
	lines := parse.Lines("cmd/day6/input.txt")
	forms := parseForms(lines)
	fmt.Printf("Part 1: %d\n", sumOfYesCounts(forms, customsForm.totalYeses))
	fmt.Printf("Part 2: %d\n", sumOfYesCounts(forms, customsForm.totalConsensusYeses))
}

func sumOfYesCounts(forms []customsForm, yesFunc func(customsForm) int) int {
	sum := 0
	for _, form := range forms {
		sum += yesFunc(form)
	}
	return sum
}

func parseForms(lines []string) []customsForm {
	var responses []string
	var forms []customsForm
	for _, line := range lines {
		if line == "" {
			forms = append(forms, customsForm{responses: responses})
			responses = nil
			continue
		}
		responses = append(responses, line)
	}
	forms = append(forms, customsForm{responses: responses})
	return forms
}

type customsForm struct {
	responses []string
}

func (f customsForm) totalYeses() int {
	questions := map[rune]bool{}
	for _, response := range f.responses {
		for _, r := range response {
			questions[r] = true
		}
	}
	return len(questions)
}

func (f customsForm) totalConsensusYeses() int {
	questions := map[rune]int{}
	for i := range f.responses {
		for _, r := range f.responses[i] {
			questions[r]++
		}
	}
	count := 0
	for _, c := range questions {
		if c == len(f.responses) {
			count++
		}
	}
	return count
}
