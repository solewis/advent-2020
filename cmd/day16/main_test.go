package main

import (
	"advent-2020/internal/parse"
	"reflect"
	"testing"
)

func TestErrorRate(t *testing.T) {
	rules, _, otherTickets := parseData(parse.String("input.test.txt"))
	answer := errorRate(rules, otherTickets)
	if answer != 71 {
		t.Errorf("Expected 71, but was %d", answer)
	}
}

func TestFieldOrder(t *testing.T) {
	rules, _, otherTickets := parseData(parse.String("input.test2.txt"))
	order := fieldOrder(rules, otherTickets)
	expected := []string{"row", "class", "seat"}
	if !reflect.DeepEqual(order, expected) {
		t.Errorf("field order was incorrect")
	}
}
