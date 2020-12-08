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

var accumulator int = 0
var instructions []instruction

func loadProgram(filename string) {
	instructions = make([]instruction, 0)
	b, _ := ioutil.ReadFile(filename)
	for _, instructionStr := range strings.Split(string(b), "\n") {
		if instructionStr == "" {
			continue
		}
		ins := strings.Split(instructionStr, " ")
		operator, _ := strconv.Atoi(string(ins[1]))
		instructions = append(instructions, instruction{string(ins[0]), operator})
	}
}

func runProgram() bool {
	iPtr := 0
	executed := make(map[int]struct{}, len(instructions))
	for {
		ins := &instructions[iPtr]

		if _, ok := executed[iPtr]; ok {
			return false
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
		if iPtr > len(instructions)-1 {
			return true
		}
	}
}

func resetProgram() {
	accumulator = 0
}

func main() {
	loadProgram("../input.txt")
	for i := 0; i < len(instructions); i++ {
		ins := &instructions[i]

		switch ins.operation {
		case "jmp":
			ins.operation = "nop"
		case "nop":
			ins.operation = "jmp"
		default:
			continue
		}

		if runProgram() {
			fmt.Println(accumulator)
			return
		}

		switch ins.operation {
		case "jmp":
			ins.operation = "nop"
		case "nop":
			ins.operation = "jmp"
		}

		resetProgram()
	}
}
