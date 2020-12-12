package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func findContig(numbers *[]int, num int) int {
	for idx1 := 0; idx1 < len(*numbers); idx1++ {
		sum := (*numbers)[idx1]
		idx2 := idx1 + 1
		for {
			sum += (*numbers)[idx2]
			if sum >= num {
				break
			}
			idx2++
		}

		if sum == num {
			sort.Ints((*numbers)[idx1:idx2])
			return (*numbers)[idx1] + (*numbers)[idx2]
		}
	}

	return 0
}

func main() {
	b, _ := ioutil.ReadFile("../input.txt")
	numbers := make([]int, 0)
	for _, numStr := range strings.Split(string(b), "\n") {
		if numStr == "" {
			continue
		}

		num, _ := strconv.Atoi(numStr)
		numbers = append(numbers, num)
	}
	fmt.Println(findContig(&numbers, 18272118))
}
