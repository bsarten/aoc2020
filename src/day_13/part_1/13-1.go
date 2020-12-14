package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func waitTime(time int, bus int) int {
	return bus - (time % bus)
}
func main() {
	busIDs := make([]int, 0)

	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	line := scanner.Text()
	earliestDeparture, _ := strconv.Atoi(line)
	scanner.Scan()
	for _, busID := range strings.Split(scanner.Text(), ",") {
		if busNum, err := strconv.Atoi(busID); err == nil {
			busIDs = append(busIDs, busNum)
		}
	}

	earliestBus := 0
	earliestWait := -1
	for _, busID := range busIDs {
		wait := busID - (earliestDeparture % busID)
		if earliestWait == -1 || wait < earliestWait {
			earliestWait = wait
			earliestBus = busID
		}
	}

	fmt.Println(earliestBus * earliestWait)
}
