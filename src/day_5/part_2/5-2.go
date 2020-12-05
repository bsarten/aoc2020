package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	nottakenmap := make(map[int]int)
	for i := 0; i < 1023; i++ {
		nottakenmap[i] = 0
	}
	takenmap := make(map[int]int)
	for scanner.Scan() {
		binaryStr := strings.Map(translateToBinary, scanner.Text())
		seatID, _ := strconv.ParseInt(binaryStr, 2, 64)
		takenmap[int(seatID)] = 0
		delete(nottakenmap, int(seatID))
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
