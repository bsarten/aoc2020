package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []int
	for scanner.Scan() {
		var newnum int
		newnum, _ = strconv.Atoi(scanner.Text())
		text = append(text, newnum)
	}

	for i1, num1 := range text {
		for i2 := i1; i2 < len(text); i2++ {
			for i3 := i2; i3 < len(text); i3++ {
				if num1+text[i2]+text[i3] == 2020 {
					fmt.Println(num1 * text[i2] * text[i3])
					return
				}
			}
		}
	}

}
