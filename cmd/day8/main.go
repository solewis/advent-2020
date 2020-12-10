package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"strings"
)

func main() {
	lines := parse.Lines("cmd/day8/input.txt")
	program := parseProgram(lines)
	acc, _ := runProgram(program)
	fmt.Printf("Part 1: %d\n", acc)
	fmt.Printf("Part 2: %d\n", accAtTermination(program))
}

type instruction struct {
	op  string
	arg int
}

func accAtTermination(program []instruction) int {
	swap := func(i int) {
		if program[i].op == "jmp" {
			program[i].op = "nop"
		} else {
			program[i].op = "jmp"
		}
	}
	for i := range program {
		if program[i].op == "jmp" || program[i].op == "nop" {
			swap(i)
			acc, terminated := runProgram(program)
			if terminated {
				return acc
			}
			swap(i) //swap back
		}
	}
	return -1
}

func runProgram(program []instruction) (acc int, terminated bool) {
	acc, insPtr := 0, 0
	instructionsRun := map[int]bool{}
	for {
		if instructionsRun[insPtr] {
			return acc, false
		}
		if insPtr == len(program) {
			return acc, true
		}
		instructionsRun[insPtr] = true
		switch program[insPtr].op {
		case "nop":
			insPtr++
		case "acc":
			acc += program[insPtr].arg
			insPtr++
		case "jmp":
			insPtr += program[insPtr].arg
		}
	}
}

func parseProgram(lines []string) []instruction {
	var instructions []instruction
	for _, line := range lines {
		parts := strings.Split(line, " ")
		instructions = append(instructions, instruction{parts[0], parse.Int(parts[1])})
	}
	return instructions
}
