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
	var highestSeatID int64 = 0
	for scanner.Scan() {
		binaryStr := strings.Map(translateToBinary, scanner.Text())
		seatID, _ := strconv.ParseInt(binaryStr, 2, 64)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	fmt.Println(highestSeatID)
}
