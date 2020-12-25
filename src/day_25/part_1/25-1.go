package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func findLoopSize(pK int) int {
	loopSize := 0
	number := 1
	for {
		loopSize++
		number = (number * 7) % 20201227
		if number == pK {
			break
		}
	}

	return loopSize
}

func findEncryptionKey(loopSize int, pK int) int {
	number := 1
	for i := 0; i < loopSize; i++ {
		number = (number * pK) % 20201227
	}

	return number
}

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	pKs := strings.Split(string(b), "\n")
	pK1, _ := strconv.Atoi(pKs[0])
	pK2, _ := strconv.Atoi(pKs[1])
	loopSize1 := findLoopSize(pK1)
	encryptionKey := findEncryptionKey(loopSize1, pK2)
	fmt.Println(encryptionKey)
}
