package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func binsearch(instructions string) int {
	high := int(math.Pow(2, float64(len(instructions))))
	low := 0
	for _, instruction := range instructions {
		if instruction == 'B' || instruction == 'R' {
			low = (high-low+1)/2 + low
		} else {
			high = (high-low+1)/2 + low - 1
		}
	}

	return low
}

func main() {
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	highestSeatID := 0
	for scanner.Scan() {
		line := scanner.Text()
		rowstr := line[0:7]
		seatstr := line[7:]

		seat := binsearch(seatstr)
		row := binsearch(rowstr)

		seatID := row*8 + seat
		if seatID > highestSeatID {
			highestSeatID = seatID
		}

	}

	fmt.Println(highestSeatID)
}
