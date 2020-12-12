package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func seesOccupied(places []string, row int, col int, rowdelta int, coldelta int) bool {
	i := row
	j := col
	for {
		i += rowdelta
		j += coldelta
		if i < 0 || i > len(places)-1 || j < 0 || j > len(places[0])-1 {
			return false
		}
		if places[i][j] == '#' {
			return true
		}
		if places[i][j] != '.' {
			return false
		}
	}
}

func countOccupied(places []string) int {
	count := 0
	for rowIdx, row := range places {
		for colIdx := range row {
			if places[rowIdx][colIdx] == '#' {
				count++
			}
		}
	}
	return count
}

func applyRules(places []string) bool {
	changed := false
	oldPlaces := make([]string, len(places))
	copy(oldPlaces, places)

	for rowIdx, row := range places {
		for colIdx := range row {
			switch oldPlaces[rowIdx][colIdx] {
			case 'L':
				if !seesOccupied(oldPlaces, rowIdx, colIdx, 0, -1) &&
					!seesOccupied(oldPlaces, rowIdx, colIdx, 0, 1) &&
					!seesOccupied(oldPlaces, rowIdx, colIdx, 1, -1) &&
					!seesOccupied(oldPlaces, rowIdx, colIdx, 1, 1) &&
					!seesOccupied(oldPlaces, rowIdx, colIdx, 1, 0) &&
					!seesOccupied(oldPlaces, rowIdx, colIdx, -1, 1) &&
					!seesOccupied(oldPlaces, rowIdx, colIdx, -1, -1) &&
					!seesOccupied(oldPlaces, rowIdx, colIdx, -1, 0) {
					places[rowIdx] = places[rowIdx][:colIdx] + string('#') + places[rowIdx][colIdx+1:]
					changed = true
				}
			case '#':
				count := 0
				if seesOccupied(oldPlaces, rowIdx, colIdx, 0, -1) {
					count++
				}
				if seesOccupied(oldPlaces, rowIdx, colIdx, 0, 1) {
					count++
				}
				if seesOccupied(oldPlaces, rowIdx, colIdx, 1, 1) {
					count++
				}
				if seesOccupied(oldPlaces, rowIdx, colIdx, 1, -1) {
					count++
				}
				if seesOccupied(oldPlaces, rowIdx, colIdx, 1, 0) {
					count++
				}
				if seesOccupied(oldPlaces, rowIdx, colIdx, -1, 1) {
					count++
				}
				if seesOccupied(oldPlaces, rowIdx, colIdx, -1, -1) {
					count++
				}
				if seesOccupied(oldPlaces, rowIdx, colIdx, -1, 0) {
					count++
				}
				if count >= 5 {
					places[rowIdx] = places[rowIdx][:colIdx] + string('L') + places[rowIdx][colIdx+1:]
					changed = true
				}
			}
		}
	}

	return changed
}

func main() {
	b, _ := ioutil.ReadFile("../input.txt")
	places := strings.Split(string(b), "\n")
	for applyRules(places[0:99]) {
	}
	fmt.Println(countOccupied(places[0:99]))
}
