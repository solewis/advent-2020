package main

import "fmt"

func main() {
	seed := []int{19, 20, 14, 0, 9, 1}
	p1 := runGame(seed, 2020)
	fmt.Printf("Part 1: %d\n", p1)
	p2 := runGame(seed, 30000000)
	fmt.Printf("Part 2: %d\n", p2)
}

func runGame(seed []int, length int) int {
	type data struct {
		turn, prevTurn int
	}
	gameState := map[int]*data{}
	spokenValue := 0
	for i, num := range seed {
		gameState[num] = &data{turn: i + 1}
		spokenValue = num
	}
	for turn := len(seed) + 1; turn <= length; turn++ {
		previousData := gameState[spokenValue]
		if previousData.prevTurn != 0 {
			spokenValue = turn - 1 - previousData.prevTurn
		} else {
			spokenValue = 0
		}
		if spokenData, exists := gameState[spokenValue]; exists {
			spokenData.prevTurn = spokenData.turn
			spokenData.turn = turn
		} else {
			gameState[spokenValue] = &data{turn: turn}
		}
	}
	return spokenValue
}
