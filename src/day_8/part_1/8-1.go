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
	runCount  int
}

var accumulator int = 0
var instructions []instruction = make([]instruction, 0)

func loadProgram(filename string) {
	b, _ := ioutil.ReadFile(filename)
	for _, instructionStr := range strings.Split(string(b), "\n") {
		if instructionStr == "" {
			continue
		}
		ins := strings.Split(instructionStr, " ")
		operator, _ := strconv.Atoi(string(ins[1]))
		instructions = append(instructions, instruction{string(ins[0]), operator, 0})
	}
}

func runProgram() {
	intPtr := 0
	for {
		ins := &instructions[intPtr]
		if ins.runCount > 0 {
			fmt.Println(accumulator)
			return
		}
		switch ins.operation {
		case "nop":
			intPtr++
		case "acc":
			accumulator += ins.argument
			intPtr++
		case "jmp":
			intPtr += ins.argument
		}
		ins.runCount++

	}
}

func main() {
	loadProgram("../input.txt")
	runProgram()
}
