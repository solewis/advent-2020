package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"strings"
)

func main() {
	data := parse.String("cmd/day19/input.txt")
	parts := strings.Split(data, "\n\n")
	rules := parseRules(strings.Split(parts[0], "\n"))
	messages := strings.Split(parts[1], "\n")
	fmt.Printf("Part 1: %d\n", rules.countMessagesMatchingRule(0, messages))
	fmt.Printf("Part 2: %d\n", p2(rules, messages))
}

func p2(ruleMap ruleMap, messages []string) int {
	// rule 8 becomes an infinite loop where any number of the original rule 8s in a row is valid
	// rule 11 becomes an infinite expansion where you can have 42 31, or 42 42 31 31, or 42 42 42 31 31 31, etc.
	// rule 0 is 8 11
	// from that we can infer that to satisfy rule 0, the message must start with n number of 42s
	//   then it must have p number of 31s. As long as there is at least 1 31
	//   and as long as there are more 42s than than 31s,
	//   and as long as the length of the 42s and 31s is as long or longer than the message it will satisfy rule 0
	rule42Len, rule31Len := 0, 0
	for k := range ruleMap[42] {
		rule42Len = len(k)
		break
	}
	for k := range ruleMap[31] {
		rule31Len = len(k)
		break
	}

	validCount := 0
	for _, message := range messages {
		num42s, num31s := 0, 0
		//start from the front of the message and count the 42s
		for i := 0; i < len(message); i+=rule42Len {
			x := message[i : i+rule42Len]
			if ruleMap[42][x] {
				num42s++
			} else {
				break
			}
		}
		//start from the back of the message and count the 31s
		for i := len(message) - 1; i >= 0; i-=rule31Len {
			x := message[i-rule31Len + 1 : i+1]
			if ruleMap[31][x] {
				num31s++
			} else {
				break
			}
		}

		if num31s > 0 && num42s > num31s && (num31s * rule31Len + num42s * rule42Len) >= len(message) {
			validCount++
		}
	}
	return validCount
}

type ruleMap map[int]rule
type rule map[string]bool

func (r ruleMap) countMessagesMatchingRule(ruleNum int, messages []string) int {
	rule := r[ruleNum]
	count := 0
	for _, message := range messages {
		if rule[message] {
			count++
		}
	}
	return count
}

func parseRules(lines []string) ruleMap {
	//rules are initially stored in a map of the rule number to the list of allowed values
	rMap := map[int][]string{}
	// loop through the rules and resolve all the ones possible
	// at first it will be all rules which are a letter
	// then all the rules which rely on those rules are solvable, keep looping until fully solved
	for {
		for _, line := range lines {
			lineParts := strings.Split(line, ": ")
			ruleNum := parse.Int(lineParts[0])
			if strings.HasPrefix(lineParts[1], "\"") {
				rMap[ruleNum] = []string{string(lineParts[1][1])}
			} else {
				ruleParts := strings.Split(lineParts[1], " | ")
				ruleSolvable := true
				for _, rulePart := range ruleParts {
					parts := strings.Split(rulePart, " ")
					for _, referencedRuleNum := range parts {
						if _, exists := rMap[parse.Int(referencedRuleNum)]; !exists {
							ruleSolvable = false
						}
					}
				}
				if ruleSolvable {
					var newRules []string
					for _, rulePart := range ruleParts {
						parts := strings.Split(rulePart, " ")
						var nums []int
						for _, ruleNumStr := range parts {
							nums = append(nums, parse.Int(ruleNumStr))
						}
						newRules = append(newRules, combineRules(nums, rMap)...)
					}
					rMap[ruleNum] = newRules
				}
			}
		}
		if len(rMap) == len(lines) {
			break
		}
	}

	//convert the list of allowed values to a set for more efficient lookup
	ruleMapWithSet := ruleMap{}
	for ruleNum, ruleList := range rMap {
		rule := map[string]bool{}
		for _, r := range ruleList {
			rule[r] = true
		}
		ruleMapWithSet[ruleNum] = rule
	}
	return ruleMapWithSet
}

//given a list of rule numbers (e.g. 4, 1, 5) will find all the allowed values for that combination of rules
//example: if 4="ab","ba" and 1="b" and 5="a", then the combos would be "abba" and "baba"
func combineRules(nums []int, ruleMap map[int][]string) []string {
	first := nums[0]
	if len(nums) == 1 {
		return ruleMap[first]
	}
	rules := ruleMap[first]
	var combos []string
	for _, rule := range rules {
		subCombos := combineRules(nums[1:], ruleMap)
		for _, subCombo := range subCombos {
			combos = append(combos, rule+subCombo)
		}
	}
	return combos
}
