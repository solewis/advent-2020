package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	lines := parse.Lines("cmd/day4/input.txt")
	passports := parsePassports(lines)
	fmt.Printf("Part 1: %d\n", countValid(passports, passport.isValid))
	fmt.Printf("Part 2: %d\n", countValid(passports, passport.isValidStrict))
}

func countValid(passports []passport, checkFunc func(passport) bool) int {
	count := 0
	for _, p := range passports {
		if checkFunc(p) {
			count++
		}
	}
	return count
}

func parsePassports(lines []string) []passport {
	var passportData []string
	var passports []passport
	for _, line := range lines {
		if line == "" {
			passports = append(passports, parsePassport(passportData))
			passportData = nil
			continue
		}
		passportData = append(passportData, strings.Split(line, " ")...)
	}
	passports = append(passports, parsePassport(passportData))
	return passports
}

func parsePassport(data []string) passport {
	p := passport{}
	for _, d := range data {
		parts := strings.Split(d, ":")
		switch parts[0] {
		case "byr":
			p.birthYear = parse.Int(parts[1])
		case "iyr":
			p.issueYear = parse.Int(parts[1])
		case "eyr":
			p.expirationYear = parse.Int(parts[1])
		case "hgt":
			p.height = parts[1]
		case "hcl":
			p.hairColor = parts[1]
		case "ecl":
			p.eyeColor = parts[1]
		case "pid":
			p.passportId = parts[1]
		case "cid":
			p.countryId = parts[1]
		}
	}
	return p
}

type passport struct {
	birthYear, issueYear, expirationYear               int
	height, hairColor, eyeColor, passportId, countryId string
}

func (p passport) isValid() bool {
	return p.birthYear != 0 &&
		p.issueYear != 0 &&
		p.expirationYear != 0 &&
		p.height != "" &&
		p.hairColor != "" &&
		p.eyeColor != "" &&
		p.passportId != ""
}

func (p passport) isValidStrict() bool {
	if !p.isValid() {
		return false
	}
	//birth year
	if p.birthYear < 1920 || p.birthYear > 2002 {
		return false
	}
	//issue year
	if p.issueYear < 2010 || p.issueYear > 2020 {
		return false
	}
	//expiration year
	if p.expirationYear < 2020 || p.expirationYear > 2030 {
		return false
	}
	//height
	if len(p.height) < 3 {
		return false
	}
	heightNum := parse.Int(p.height[:len(p.height)-2])
	switch {
	case strings.HasSuffix(p.height, "cm"):
		if heightNum < 150 || heightNum > 193 {
			return false
		}
	case strings.HasSuffix(p.height, "in"):
		if heightNum < 59 || heightNum > 76 {
			return false
		}
	default: return false
	}
	//hair color
	hairColorRegex := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	if !hairColorRegex.MatchString(p.hairColor) {
		return false
	}
	//eye color
	eyeColorRegex := regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`)
	if !eyeColorRegex.MatchString(p.eyeColor) {
		return false
	}
	//passport id
	passportIdRegex := regexp.MustCompile(`^[0-9]{9}$`)
	if !passportIdRegex.MatchString(p.passportId) {
		return false
	}

	return true
}
