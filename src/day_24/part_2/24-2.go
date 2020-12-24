package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type coordinates struct {
	x int
	y int
}

type coordinatesMap map[coordinates]empty

type empty struct{}

var listOfNeighborDirections []string = []string{"e", "se", "sw", "w", "nw", "ne"}

func coordinatesFrom(coords coordinates, directions string) coordinates {
	re := regexp.MustCompile(`(e|se|sw|w|nw|ne)`)
	match := re.FindAllString(directions, -1)

	for i := 0; i < len(match); i++ {
		direction := match[i]
		coords = coordinatesOfNeighbor(coords, direction)
	}

	return coords
}

func coordinatesOfNeighbor(coords coordinates, direction string) coordinates {
	switch direction {
	case "e":
		coords.x++
	case "w":
		coords.x--
	case "se":
		coords.x++
		coords.y--
	case "sw":
		coords.y--
	case "nw":
		coords.x--
		coords.y++
	case "ne":
		coords.y++
	}

	return coords
}

func followDirections(inputFile string) coordinatesMap {
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	blackTiles := make(coordinatesMap)

	for _, line := range strings.Split(string(b), "\n") {
		coords := coordinates{0, 0}
		if line == "" {
			continue
		}

		coords = coordinatesFrom(coords, line)

		if _, exists := blackTiles[coords]; exists {
			delete(blackTiles, coords)
		} else {
			blackTiles[coords] = empty{}
		}
	}

	return blackTiles
}

func countBlackNeighbors(blackTiles coordinatesMap, coords coordinates) int {
	count := 0
	for _, direction := range listOfNeighborDirections {
		neighborCoords := coordinatesOfNeighbor(coords, direction)
		if _, exists := blackTiles[neighborCoords]; exists {
			count++
		}
	}

	return count
}

func simulateDay(blackTiles coordinatesMap) coordinatesMap {
	destTiles := make(map[coordinates]empty, 0)
	visited := make(map[coordinates]empty, 0)

	for coords := range blackTiles {
		visited[coords] = empty{}
		numBlackNeighbors := countBlackNeighbors(blackTiles, coords)
		if numBlackNeighbors > 0 && numBlackNeighbors <= 2 {
			destTiles[coords] = empty{}
		}

		// white neighbors
		for _, direction := range listOfNeighborDirections {
			neighborCoords := coordinatesOfNeighbor(coords, direction)
			if _, exists := blackTiles[neighborCoords]; !exists {
				if _, exists := visited[neighborCoords]; !exists {
					if countBlackNeighbors(blackTiles, neighborCoords) == 2 {
						destTiles[neighborCoords] = empty{}
					}
				}
			}
			visited[neighborCoords] = empty{}
		}
	}

	return destTiles
}

func main() {
	blackTiles := followDirections("../input.txt")
	for i := 0; i < 100; i++ {
		blackTiles = simulateDay(blackTiles)
	}
	fmt.Println(len(blackTiles))
}
