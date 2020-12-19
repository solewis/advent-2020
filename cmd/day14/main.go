package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	maskRegex = regexp.MustCompile(`^mask = ([01X]{36})$`)
	memRegex = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
)

func main() {
	instructions := parse.Lines("cmd/day14/input.txt")
	fmt.Printf("Part 1: %d\n", sumOfFinalMemory(instructions, runProgram))
	fmt.Printf("Part 2: %d\n", sumOfFinalMemory(instructions, runProgramV2))
}

func sumOfFinalMemory(instructions []string, runFunc func([]string)map[uint64]uint64) int {
	memory := runFunc(instructions)
	count := 0
	for _, v := range memory {
		count += int(v)
	}
	return count
}

func runProgram(instructions []string) map[uint64]uint64 {
	memory := map[uint64]uint64{}
	var zeroMask, oneMask uint64
	for _, ins := range instructions {
		if matches := maskRegex.FindStringSubmatch(ins); matches != nil {
			mask := matches[1]
			zeroMask, _ = strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 64)
			oneMask, _ = strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 64)
		} else {
			matches := memRegex.FindStringSubmatch(ins)
			addr, _ := strconv.ParseUint(matches[1], 10, 64)
			val, _ := strconv.ParseUint(matches[2], 10, 64)
			memory[addr] = (val & zeroMask) | oneMask
		}
	}
	return memory
}

func runProgramV2(instructions []string) map[uint64]uint64 {
	memory := map[uint64]uint64{}
	var mask string
	var oneMask uint64
	for _, ins := range instructions {
		if matches := maskRegex.FindStringSubmatch(ins); matches != nil {
			mask = matches[1]
			oneMask, _ = strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 64)
		} else {
			matches := memRegex.FindStringSubmatch(ins)
			val, _ := strconv.ParseUint(matches[2], 10, 64)

			addrStr := fmt.Sprintf("%036b", parse.Int(matches[1]))
			numFloating := strings.Count(mask, "X")
			numCombos := int(math.Pow(2, float64(numFloating)))
			// count from 0 to {numCombos} and use "i"'s binary representation to replace the x's
			// e.g. 4 combos would result in binary values 00, 01, 10, and 11 (decimal 0,1,2,3).
			// Replace the x's with those to get each address location
			for i := 0; i < numCombos; i++ {
				x := 0
				b := fmt.Sprintf("%0*b", numFloating, i)
				addrCombo := make([]byte, len(addrStr))
				for j := range mask {
					if mask[j] == 'X' {
						addrCombo[j] = b[x]
						x++
					} else {
						addrCombo[j] = addrStr[j]
					}
				}
				addr, _ := strconv.ParseUint(string(addrCombo), 2, 64)
				memory[addr | oneMask] = val
			}
		}
	}
	return memory
}
