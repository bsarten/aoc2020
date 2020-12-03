package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("../input.txt")
	var low_occurance, high_occurance int
	var letter, password string
	var number_valid int
	for {
		_, err := fmt.Fscanf(file, "%d-%d %1s: %s\n", &low_occurance, &high_occurance, &letter, &password)
		if err == io.EOF {
			break
		}

		actual_occurance := strings.Count(password, letter)
		if actual_occurance >= low_occurance && actual_occurance <= high_occurance {
			number_valid++
		}
	}

	//	q := new(Q)
	//	q.Init()

	fmt.Println(number_valid)
}
