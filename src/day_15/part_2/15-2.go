package main

import (
	"fmt"
)

type ElfNumber struct {
	order       int
	timesSpoken int
}

func main() {
	numbers := map[int]ElfNumber{
		2:  {1, 1},
		0:  {2, 1},
		1:  {3, 1},
		9:  {4, 1},
		5:  {5, 1},
		19: {6, 1},
	}

	lastNum := 19
	for turn := len(numbers) + 1; turn <= 30000000; turn++ {
		last, _ := numbers[lastNum]
		if last.timesSpoken == 1 {
			lastNum = 0
		} else {
			numbers[lastNum] = ElfNumber{turn - 1, last.timesSpoken}
			lastNum = turn - 1 - last.order
		}

		if num, ok := numbers[lastNum]; ok {
			numbers[lastNum] = ElfNumber{num.order, num.timesSpoken + 1}
		} else {
			numbers[lastNum] = ElfNumber{turn, 1}
		}
	}

	fmt.Println(lastNum)
}
