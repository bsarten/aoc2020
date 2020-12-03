package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, _ := os.Open("../input.txt")
	var position1, position2 int
	var letter, password string
	var numberValid int
	for {
		_, err := fmt.Fscanf(file, "%d-%d %1s: %s", &position1, &position2, &letter, &password)
		if err == io.EOF {
			break
		}

		char1 := password[position1-1 : position1]
		char2 := password[position2-1 : position2]

		if (char1 == letter && char2 != letter) || (char2 == letter && char1 != letter) {
			numberValid++
		}
	}

	fmt.Println(numberValid)
}
