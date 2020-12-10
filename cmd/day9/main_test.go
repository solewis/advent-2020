package main

import (
	"advent-2020/internal/parse"
	"testing"
)

func TestFindFirstInvalid(t *testing.T) {
	data := parse.Ints("input.test.txt", "\n")
	answer := findFirstInvalid(data, 5)
	if answer != 127 {
		t.Errorf("Find first invalid failed, expected 127 but was %d", answer)
	}
}

func TestFindEncryptionWeakness(t *testing.T) {
	data := parse.Ints("input.test.txt", "\n")
	firstInvalid := findFirstInvalid(data, 5)
	r := findEncryptionWeakness(firstInvalid, data)
	if r != 62 {
		t.Errorf("Error finding range, expected 62, but was %d", r)
	}
}
