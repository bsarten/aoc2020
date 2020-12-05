package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set"
)

func translateToBinary(r rune) rune {
	switch r {
	case 'B', 'R':
		return '1'
	default:
		return '0'
	}
}

func main() {
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	allSeats := mapset.NewSet()
	for i := 0; i < 1023; i++ {
		allSeats.Add(i)
	}
	seatsTaken := mapset.NewSet()
	for scanner.Scan() {
		binaryStr := strings.Map(translateToBinary, scanner.Text())
		seatID, _ := strconv.ParseInt(binaryStr, 2, 64)
		seatsTaken.Add(int(seatID))
	}

	seatsNotTaken := allSeats.Difference(seatsTaken)
	for seat := range seatsNotTaken.Iter() {
		if seatsTaken.Contains(seat.(int)-1) && seatsTaken.Contains(seat.(int)+1) {
			fmt.Println(seat)
		}
	}
}
