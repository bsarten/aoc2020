package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

var alreadyCounted map[int]int = make(map[int]int, 0)

func countArrangements(voltages *[]int, idx int) int {
	if idx == len(*voltages)-1 {
		return 1
	}

	if counted, ok := alreadyCounted[idx]; ok {
		return counted
	}

	count := 0
	for idxplus := 1; idxplus <= 3; idxplus++ {
		if (idx + idxplus) > (len(*voltages) - 1) {
			break
		}
		diff := (*voltages)[idxplus+idx] - (*voltages)[idx]
		if diff <= 3 {
			count += countArrangements(voltages, idx+idxplus)
		}
	}

	alreadyCounted[idx] = count
	return count
}

func main() {
	b, _ := ioutil.ReadFile("../input.txt")
	voltages := make([]int, 0)
	voltages = append(voltages, 0)
	for _, numStr := range strings.Split(string(b), "\n") {
		if numStr == "" {
			continue
		}

		num, _ := strconv.Atoi(numStr)
		voltages = append(voltages, num)
	}
	sort.Ints(voltages)
	fmt.Println(countArrangements(&voltages, 0))
}
