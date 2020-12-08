package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	argument  int
}

func loadProgram(filename string) *[]instruction {
	instructions := make([]instruction, 0)
	b, _ := ioutil.ReadFile(filename)
	for _, instructionStr := range strings.Split(string(b), "\n") {
		if instructionStr == "" {
			continue
		}
		ins := strings.Split(instructionStr, " ")
		operator, _ := strconv.Atoi(string(ins[1]))
		instructions = append(instructions, instruction{string(ins[0]), operator})
	}
	return &instructions
}

func runProgram(instructions *[]instruction) (bool, int) {
	iPtr := 0
	accumulator := 0
	executed := make(map[int]struct{}, len(*instructions))
	for {
		ins := (*instructions)[iPtr]

		if _, ok := executed[iPtr]; ok {
			return false, accumulator
		}
		executed[iPtr] = struct{}{}

		switch ins.operation {
		case "nop":
			iPtr++
		case "acc":
			accumulator += ins.argument
			iPtr++
		case "jmp":
			iPtr += ins.argument
		}
		if iPtr > len(*instructions)-1 {
			return true, accumulator
		}
	}
}

func main() {
	instructions := loadProgram("../input.txt")
	for idx := range *instructions {
		ins := &(*instructions)[idx]

		switch ins.operation {
		case "jmp":
			ins.operation = "nop"
		case "nop":
			ins.operation = "jmp"
		default:
			continue
		}

		if ok, accumulator := runProgram(instructions); ok {
			fmt.Println(accumulator)
			return
		}

		switch ins.operation {
		case "jmp":
			ins.operation = "nop"
		case "nop":
			ins.operation = "jmp"
		}
	}
}
