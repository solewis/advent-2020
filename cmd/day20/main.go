package main

import (
	"advent-2020/internal/parse"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	data := parse.String("cmd/day20/input.txt")
	fmt.Printf("Part 1: %d\n", p1(data))
}

func p1(data string) int {
	tiles := parseTiles(data)
	edges := parseEdges(tiles)

	//create a map of each edge to the other edge it matches. Edges with no match will not be in the map
	edgeMatches := map[edgeKey]edgeKey{}
	for key, edge := range edges {
		if _, exists := edgeMatches[key]; exists {
			continue
		}
		for key2, edge2 := range edges {
			if key == key2 {
				continue
			}
			if edge == edge2 || reverse(edge) == edge2 {
				if _, exists := edgeMatches[key2]; exists {
					panic("multiple matches for an edge")
				}
				edgeMatches[key] = key2
				edgeMatches[key2] = key
			}
		}
	}

	//find corner edges
	//for each edge with no matches, if the edge adjacent in one direction also has no matches, it is a corner
	result := 1
	for key := range edges {
		if _, hasMatch := edgeMatches[key]; !hasMatch {
			if _, adjHasMatch := edgeMatches[edgeKey{key.id, (key.edge + 1)%4}]; !adjHasMatch {
				result *= key.id
			}
		}
	}
	return result

	//tiles := parseTiles(lines)
	//size := int(math.Sqrt(float64(len(tiles))))
	//initialGrid := make([][]tile, size)
	//finalGrid, _ := calc(tiles, initialGrid, 0)
	//return finalGrid[0][0].id * finalGrid[0][size-1].id * finalGrid[size-1][0].id * finalGrid[size-1][size-1].id
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type tile []string

var idRegex = regexp.MustCompile(`^Tile (\d+):$`)

func parseTiles(data string) map[int]tile {
	tiles := map[int]tile{}
	parts := strings.Split(data, "\n\n")
	for _, tileData := range parts {
		tileLines := strings.Split(tileData, "\n")
		matches := idRegex.FindStringSubmatch(tileLines[0])
		id := parse.Int(matches[1])
		tiles[id] = tileLines[1:]
	}
	return tiles
}

type edgeKey struct {
	id, edge int
}

func parseEdges(tiles map[int]tile) map[edgeKey]string {
	edges := map[edgeKey]string{}
	for id, tile := range tiles {

		var right, left []byte
		for _, tileRow := range tile {
			right = append(right, tileRow[len(tileRow)-1])
			left = append(left, tileRow[0])
		}

		edges[edgeKey{id, 0}] = tile[0]           //top
		edges[edgeKey{id, 1}] = string(right)     //right
		edges[edgeKey{id, 2}] = tile[len(tile)-1] //bottom
		edges[edgeKey{id, 3}] = string(left)      //left
	}
	return edges
}

//func parseTiles(lines []string) []tile {
//	return nil
//}
//
//type tile struct {
//	id    int
//	image []string
//}
//
//func (t tile) orientations() []tile {
//	var orientations []tile
//	flipHorizontal := make([][]rune, len(t.image))
//	for rowIndex, row := range t.image {
//		flipHorizontal[rowIndex] = make([]rune, len(row))
//		for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
//			flipHorizontal[rowIndex][i], flipHorizontal[rowIndex][j] = row[j], row[i]
//		}
//	}
//
//	flipVertical := make([][]rune, len(t.image))
//	for i, j := 0, len(t.image)-1; i < j; i, j = i+1, j-1 {
//		flipVertical[i], flipVertical[j] = t.image[j], t.image[i]
//	}
//
//	return nil
//}
//
//func allOrientations(tiles []tile, acc map[tile]bool) []tile {
//	if tiles == nil {
//		return acc
//	}
//
//	var orientations []tile
//	flipHorizontal := make([][]rune, len(t.image))
//	for rowIndex, row := range t.image {
//		flipHorizontal[rowIndex] = make([]rune, len(row))
//		for i, j := 0, len(row)-1; i < j; i, j = i+1, j-1 {
//			flipHorizontal[rowIndex][i], flipHorizontal[rowIndex][j] = row[j], row[i]
//		}
//	}
//
//	flipVertical := make([][]rune, len(t.image))
//	for i, j := 0, len(t.image)-1; i < j; i, j = i+1, j-1 {
//		flipVertical[i], flipVertical[j] = t.image[j], t.image[i]
//	}
//
//	return nil
//}

//type grid [][]tile
//
//func (g grid) update(position int, tile tile) {
//
//}
//
//func (g grid) valid() bool {
//	return false
//}
//
//func isValid() bool {
//	return false
//}
//
//func calc(tiles []tile, gridSoFar grid, position int) (grid, bool) {
//	if len(tiles) == 0 {
//		return gridSoFar, true
//	}
//	for i, tile := range tiles {
//		for _, orientation := range tile.orientations() {
//			gridSoFar.update(position, orientation)
//			if gridSoFar.valid() {
//				if finalGrid, ok := calc(append(tiles[:i], tiles[i+1:]...), gridSoFar, position+1); ok {
//					return finalGrid, true
//				}
//			}
//		}
//	}
//	panic("no solution")
//}
