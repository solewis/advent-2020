package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	data := parse.String("cmd/day16/input.txt")
	rules, myTicket, otherTickets := parseData(data)
	fmt.Printf("Part 1: %d\n", errorRate(rules, otherTickets))

	order := fieldOrder(rules, otherTickets)
	fmt.Printf("Part 2: %d\n", part2(order, myTicket))
}

var (
	ruleRegex = regexp.MustCompile(`^(.+): (\d+)-(\d+) or (\d+)-(\d+)$`)
)

type rules map[string]map[int]bool

func part2(order []string, myTicket []int) int {
	answer := 1
	for i, field := range order {
		if strings.HasPrefix(field, "departure") {
			answer *= myTicket[i]
		}
	}
	return answer
}

func fieldOrder(rules rules, tickets [][]int) []string {
	valid := validTickets(rules, tickets)
	//start with map of positions to all the fields valid at that position (every field to start)
	positionMap := map[int]map[string]bool{}
	for i := 0; i < len(rules); i++ {
		allPositions := map[string]bool{}
		for k := range rules {
			allPositions[k] = true
		}
		positionMap[i] = allPositions
	}

	//go through all the tickets and pare down the valid positions for each field
	for _, ticket := range valid {
		for position, val := range ticket {
			for field, allowedVals := range rules {
				if !allowedVals[val] {
					positionMap[position][field] = false
				}
			}
		}
	}

	//build out the order
	//all positions with only 1 valid field become that field. Remove that field from the other positions, repeat until all positions hold only 1 field
	order := make([]string, len(positionMap))
	found := 0
	for {
		for position, possibleFields := range positionMap {
			var possibleFieldsList []string
			for field, valid := range possibleFields {
				if valid {
					possibleFieldsList = append(possibleFieldsList, field)
				}
			}
			if len(possibleFieldsList) == 1 {
				order[position] = possibleFieldsList[0]
				found++
				if found == len(order) {
					return order
				}
				//remove field from other positions
				for _, possibleFields := range positionMap {
					possibleFields[possibleFieldsList[0]] = false
				}
			}
		}
	}
}

func validTickets(rules rules, tickets [][]int) [][]int {
	allAllowedValues := map[int]bool{}
	for _, v := range rules {
		for k := range v {
			allAllowedValues[k] = true
		}
	}
	var validTickets [][]int
	for _, ticket := range tickets {
		valid := true
		for _, val := range ticket {
			if !allAllowedValues[val] {
				valid = false
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}
	return validTickets
}

func errorRate(rules rules, tickets [][]int) int {
	allAllowedValues := map[int]bool{}
	// compile list of every allowed value in all fields
	for _, v := range rules {
		for k := range v {
			allAllowedValues[k] = true
		}
	}
	//sum of all fields in any ticket which aren't valid anywhere
	sumOfInvalid := 0
	for _, ticket := range tickets {
		for _, val := range ticket {
			if !allAllowedValues[val] {
				sumOfInvalid += val
			}
		}
	}
	return sumOfInvalid
}

func parseData(data string) (rules, []int, [][]int) {
	rules := rules{}
	var myTicket []int
	var otherTickets [][]int
	dataParts := strings.Split(data, "\n\n")

	//parse rules
	for _, rule := range strings.Split(dataParts[0], "\n") {
		matches := ruleRegex.FindStringSubmatch(rule)
		field := matches[1]
		allowedVals := map[int]bool{}
		for i := parse.Int(matches[2]); i <= parse.Int(matches[3]); i++ {
			allowedVals[i] = true
		}
		for i := parse.Int(matches[4]); i <= parse.Int(matches[5]); i++ {
			allowedVals[i] = true
		}
		rules[field] = allowedVals
	}

	//parse my ticket
	myTicketParts := strings.Split(dataParts[1], "\n")
	for _, val := range strings.Split(myTicketParts[1], ",") {
		myTicket = append(myTicket, parse.Int(val))
	}

	//parse other tickets
	otherTicketParts := strings.Split(dataParts[2], "\n")
	for _, ticketStr := range otherTicketParts[1:] {
		var ticket []int
		for _, val := range strings.Split(ticketStr, ",") {
			ticket = append(ticket, parse.Int(val))
		}
		otherTickets = append(otherTickets, ticket)
	}

	return rules, myTicket, otherTickets
}
