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

func main() {
	blackTiles := followDirections("../input.txt")
	fmt.Println(len(blackTiles))
}
