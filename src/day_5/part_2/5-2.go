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
	nottakenmap := make(map[int]int)
	for i := 0; i < 1023; i++ {
		nottakenmap[i] = 0
	}
	takenmap := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		rowstr := line[0:7]
		seatstr := line[7:]

		seat := binsearch(seatstr)
		row := binsearch(rowstr)

		seatID := row*8 + seat
		takenmap[seatID] = 0
		delete(nottakenmap, seatID)
	}

	for seat := range nottakenmap {
		_, ok := takenmap[seat-1]
		if ok {
			_, ok := takenmap[seat+1]
			if ok {
				fmt.Println(seat)
			}
		}
	}
}
