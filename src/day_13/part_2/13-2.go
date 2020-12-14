package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bus struct {
	index  int
	busnum int
}

func main() {
	busIDs := make([]int, 0)

	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	scanner.Scan()
	for _, busID := range strings.Split(scanner.Text(), ",") {
		if busNum, err := strconv.Atoi(busID); err == nil {
			busIDs = append(busIDs, busNum)
		} else {
			busIDs = append(busIDs, 0)
		}
	}

	time := busIDs[0]
	for busIdx, busID := range busIDs[1:] {
		if busID == 0 {
			continue
		}
		productOfBusIDs := 1
		// get product of previous buses
		for i := 0; i < busIdx+1; i++ {
			if busIDs[i] != 0 {
				productOfBusIDs *= busIDs[i]
			}
		}

		// find match, incrementing by previous product
		// to keep it a multiple of previous buses. note, this
		// only works if the input is all primes
		for (time+busIdx+1)%busID != 0 {
			time += productOfBusIDs
		}
	}
	fmt.Println(time)
}
