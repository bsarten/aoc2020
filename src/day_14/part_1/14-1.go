package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const memInstruction = 0
const maskInstruction = 1

type instruction struct {
	instructionType int
	location        int
	value           string
}

func readProgram(filename string) []instruction {
	instructions := make([]instruction, 0)

	b, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(string(b), "\n") {
		if line == "" {
			continue
		}
		eqsplit := strings.Split(line, " = ")
		re := regexp.MustCompile(`^mem\[(\d+)\]$`)
		match := re.FindStringSubmatch(eqsplit[0])
		if len(match) == 2 {
			location, _ := strconv.Atoi(match[1])
			instructions = append(instructions, instruction{memInstruction, location, eqsplit[1]})
		} else {
			instructions = append(instructions, instruction{maskInstruction, 0, eqsplit[1]})
		}
	}

	return instructions
}

func runProgram(program []instruction, memory map[int]uint64) {
	var maskZero uint64 = 0
	var maskOne uint64 = 0
	for _, instruction := range program {
		switch instruction.instructionType {
		case maskInstruction:
			var temp uint64
			temp, _ = strconv.ParseUint(strings.ReplaceAll(instruction.value, "X", "1"), 2, 64)
			maskZero = 0xFFFFFFF000000000 | temp
			temp, _ = strconv.ParseUint(strings.ReplaceAll(instruction.value, "X", "0"), 2, 64)
			maskOne = temp
		case memInstruction:
			var temp int
			temp, _ = strconv.Atoi(instruction.value)
			memory[instruction.location] = (uint64(temp) | maskOne) & maskZero
		}
	}
}

func main() {
	program := readProgram("../input.txt")
	memory := make(map[int]uint64, 0)
	runProgram(program, memory)
	var sum uint64 = 0
	for _, v := range memory {
		sum += v
	}
	fmt.Println(sum)
}