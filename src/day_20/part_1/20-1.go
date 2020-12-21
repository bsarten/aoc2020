package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type tile struct {
	id           int
	image        []string
	leftBorder   string
	rightBorder  string
	topBorder    string
	bottomBorder string
}

const PLACED_SIZE = 17

type empty struct{}

func getImageColumn(tileImage []string, col int) string {
	colStr := ""
	for i := 0; i < len(tileImage); i++ {
		colStr += string(tileImage[i][col])
	}
	return colStr
}

func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

func newTile(tileName string, tileImage []string) tile {
	tile := tile{}
	id, _ := strconv.Atoi(tileName)
	tile.id = id
	tile.image = tileImage
	tile.leftBorder = getImageColumn(tileImage, 0)
	tile.rightBorder = getImageColumn(tileImage, len(tileImage)-1)
	tile.bottomBorder = tileImage[len(tileImage)-1]
	tile.topBorder = tileImage[0]
	return tile
}

func rotateTile(tile tile) tile {
	temp := tile.topBorder
	tile.topBorder = Reverse(tile.leftBorder)
	tile.leftBorder = tile.bottomBorder
	tile.bottomBorder = Reverse(tile.rightBorder)
	tile.rightBorder = temp
	return tile
}

func flipTile(tile tile) tile {
	tile.topBorder = Reverse(tile.topBorder)
	tile.bottomBorder = Reverse(tile.bottomBorder)
	temp := tile.leftBorder
	tile.leftBorder = tile.rightBorder
	tile.rightBorder = temp
	return tile
}

func placeTiles(x int, y int, tiles []tile, usedTiles map[int]empty, placedTiles [][]tile) bool {
	if y >= PLACED_SIZE {
		return true
	}

	if usedTiles == nil {
		usedTiles = make(map[int]empty, 0)
	}
	for tileIdx := 0; tileIdx < len(tiles); tileIdx++ {
		if _, exists := usedTiles[tileIdx]; exists {
			continue
		}

		checkTile := tiles[tileIdx]
		for i := 0; i < 2; i++ {
			for j := 0; j < 4; j++ {
				if (x == 0 || checkTile.leftBorder == placedTiles[x-1][y].rightBorder) && (y == 0 || checkTile.topBorder == placedTiles[x][y-1].bottomBorder) {
					placedTiles[x][y] = checkTile
					usedTiles[tileIdx] = empty{}
					nextX := (x + 1) % PLACED_SIZE
					nextY := y + (x+1)/PLACED_SIZE
					if placeTiles(nextX, nextY, tiles, usedTiles, placedTiles) {
						return true
					} else {
						delete(usedTiles, tileIdx)
					}
				}
				checkTile = rotateTile(checkTile)
			}
			checkTile = flipTile(checkTile)
		}
	}

	return false
}

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	tilesSplit := strings.Split(string(b), "\n\n")
	tiles := make([]tile, 0)
	for _, tile := range tilesSplit {
		if tile == "" {
			continue
		}
		tileLines := strings.Split(tile, "\n")
		tileImage := tileLines[1:]
		tileDesc := tileLines[0]
		re := regexp.MustCompile(`^Tile (\d+):$`)
		match := re.FindStringSubmatch(tileDesc)
		tileName := match[1]
		newTile := newTile(tileName, tileImage)
		tiles = append(tiles, newTile)
	}

	placedTiles := make([][]tile, PLACED_SIZE)
	for i := 0; i < PLACED_SIZE; i++ {
		placedTiles[i] = make([]tile, PLACED_SIZE)
	}
	if placeTiles(0, 0, tiles, nil, placedTiles) {
		fmt.Println(placedTiles[0][0].id * placedTiles[0][PLACED_SIZE-1].id * placedTiles[PLACED_SIZE-1][0].id * placedTiles[PLACED_SIZE-1][PLACED_SIZE-1].id)
	} else {
		fmt.Println("nope")
	}
}
