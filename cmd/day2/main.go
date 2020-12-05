package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"regexp"
)

func main() {
	lines := parse.Lines("cmd/day2/input.txt")
	passwords := parsePasswords(lines)

	fmt.Printf("Part 1: %d\n", countValid(passwords, passwordFile.isValidSledRentalRules))
	fmt.Printf("Part 2: %d\n", countValid(passwords, passwordFile.isValidTobogganRules))
}

func countValid(passwords []passwordFile, ruleFunc func(passwordFile) bool) int {
	count := 0
	for _, p := range passwords {
		if ruleFunc(p) {
			count++
		}
	}
	return count
}

func parsePasswords(lines []string) []passwordFile {
	re := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): (\w+)`)
	var passwords []passwordFile
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		passwords = append(passwords, passwordFile{
			ltr:   matches[3],
			pwStr: matches[4],
			a:     parse.Int(matches[1]),
			b:     parse.Int(matches[2]),
		})
	}
	return passwords
}

type passwordFile struct {
	ltr, pwStr string
	a, b       int
}

func (p passwordFile) isValidSledRentalRules() bool {
	count := 0
	for _, r := range p.pwStr {
		if string(r) == p.ltr {
			count++
		}
	}
	return count <= p.b && count >= p.a
}

func (p passwordFile) isValidTobogganRules() bool {
	return (string(p.pwStr[p.a-1]) == p.ltr) != (string(p.pwStr[p.b-1]) == p.ltr)
}
