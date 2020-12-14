package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"math"
	"strings"
)

func main() {
	lines := parse.Lines("cmd/day13/input.txt")
	timestamp := parse.Int(lines[0])
	var busses []bus
	for i, id := range strings.Split(lines[1], ",") {
		if id != "x" {
			busses = append(busses, bus{i, parse.Int(id)})
		}
	}
	fmt.Printf("Part 1: %d\n", part1(timestamp, busses))
	fmt.Printf("Part 2: %d\n", part2(busses))
	//test()
}

type bus struct{ index, id int }

func part1(timestamp int, busses []bus) int {
	minWaitTime, minId := math.MaxInt64, 0
	for _, bus := range busses {
		waitTime := bus.id - timestamp%bus.id
		if waitTime < minWaitTime {
			minWaitTime = waitTime
			minId = bus.id
		}
	}
	return minWaitTime * minId
}

//**how I solved part 2 initially

//the brute force algorithm just looped incrementing by the id of the first bus, which is too slow
//we can align the first n busses (4 in this case) separately though, and see how often they align
//then we can run our brute force algorithm, but we start with the first time the first 4 aligned, and increment by how often they align
//until all the busses align. This reduces the number of loops needed to solve part 2 in a few seconds

//func part2(busses []bus) int {
//	bus1, bus2 := busses[0], busses[0] //index1 will be index with highest number, index2 will be second highest
//	for _, bus := range busses {
//		if bus.id > bus1.id {
//			bus2 = bus1 //first becomes second
//			bus1 = bus      //new max
//		} else if bus.id > bus2.id {
//			bus2 = bus //new second max
//		}
//	}
//	// firstAlign will be the first time the first n numbers align,
//	// secondAlign will be the second time they align
//	// aligning the first 4 busses made it fast enough
//	firstAlign, secondAlign := 0, 0
//	for id := busses[0].id; ; id += busses[0].id {
//		bus2Ok := (id+busses[1].index) %busses[1].id == 0
//		bus3Ok := (id+busses[2].index) %busses[2].id == 0
//		bus4Ok := (id+busses[3].index) %busses[3].id == 0
//		if bus2Ok && bus3Ok && bus4Ok {
//			if firstAlign == 0 {
//				firstAlign = id
//			}else {
//				secondAlign = id
//				break
//			}
//		}
//	}
//	// start at the first time the first n busses aligned and increase by how often the align
//	// until a timestamp is found which aligns all the busses
//	for currentTs := firstAlign; true; currentTs += secondAlign - firstAlign {
//		fitsPattern := true
//		for _, bus := range busses {
//			if  (currentTs+bus.index)%bus.id != 0 {
//				fitsPattern = false
//				break
//			}
//		}
//		if fitsPattern {
//			return currentTs
//		}
//	}
//	return -1
//}

// cleaner solve
// start by incrementing the timestamp along the first bus timings
// align the first bus with the second bus and find out how often they sync
// start incrementing by this new higher number (how often bus 1 and 2 sync)
// align now with the third bus. Once it aligns you can start incrementing by how often the first 3 busses align.
// continue until you get through the final bus. Solves instantly.
func part2(busses []bus) int {
	increment := busses[0].id
	currentBusIndex := 1 //bus we are aligning
	firstAlignment := 0
	for currentTs := busses[0].id; ; currentTs += increment {
		currentBus := busses[currentBusIndex]
		if (currentTs+currentBus.index)%currentBus.id == 0 {
			if currentBusIndex == len(busses)-1 {
				//all busses aligned
				return currentTs
			} else if firstAlignment == 0 {
				//first time the bus aligns
				//must find the first two times the busses align in order to see the frequency of their alignments
				firstAlignment = currentTs
			} else {
				//second time the bus aligns
				increment = currentTs - firstAlignment
				firstAlignment = 0
				currentBusIndex++
			}
		}
	}
}
