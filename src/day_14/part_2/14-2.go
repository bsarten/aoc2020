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
	location        string
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
			instructions = append(instructions, instruction{memInstruction, match[1], eqsplit[1]})
		} else {
			instructions = append(instructions, instruction{maskInstruction, "", eqsplit[1]})
		}
	}

	return instructions
}

func setMem(memory map[string]uint64, location []byte, value string, mask string, maskPosition int) {
	for {
		if maskPosition == len(location) {
			break
		}
		switch mask[maskPosition] {
		case '1':
			fmt.Println(string(location))
			location[maskPosition] = '1'
			fmt.Println(string(location))
		case 'X':
			location[maskPosition] = '0'
			setMem(memory, location, value, mask, maskPosition+1)
			location[maskPosition] = '1'
		}
		maskPosition++
	}

	temp2, _ := strconv.Atoi(value)
	memory[string(location)] = uint64(temp2)
}

func runProgram(program []instruction, memory map[string]uint64) {
	var mask string = ""
	for _, instruction := range program {
		switch instruction.instructionType {
		case maskInstruction:
			mask = instruction.value
		case memInstruction:
			temp, _ := strconv.Atoi(instruction.location)
			location := []byte(fmt.Sprintf("%036b", uint64(temp)))
			setMem(memory, location, instruction.value, mask, 0)
		}
	}
}

func main() {
	program := readProgram("../input.txt")
	memory := make(map[string]uint64, 0)
	runProgram(program, memory)
	var sum uint64 = 0
	for _, v := range memory {
		sum += v
	}
	fmt.Println(sum)
}
