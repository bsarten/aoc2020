package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type instruction struct {
	dir string
	num int
}

func normalizeHeading(heading int) int {
	if heading < 0 {
		heading += 360
	} else if heading >= 360 {
		heading -= 360
	}
	return heading
}

func move(heading int, num int, east *int, north *int) {
	switch heading {
	case 0:
		*north += num
	case 90:
		*east += num
	case 180:
		*north -= num
	case 270:
		*east -= num
	}
}

func rotate(east *int, north *int, degrees int) {
	switch degrees {
	case 90:
		temp := *east
		*east = *north
		*north = -temp
	case 180:
		*east = -*east
		*north = -*north
	case 270:
		temp := *east
		*east = -*north
		*north = temp
	}
}

func getFinalPosition(instructions []instruction) (int, int) {
	wpeast := 10
	wpnorth := 1
	east := 0
	north := 0

	for _, instruction := range instructions {
		switch instruction.dir {
		case "L":
			rotate(&wpeast, &wpnorth, 360-instruction.num)
		case "R":
			rotate(&wpeast, &wpnorth, instruction.num)
		case "F":
			east += instruction.num * wpeast
			north += instruction.num * wpnorth
		case "N":
			wpnorth += instruction.num
		case "S":
			wpnorth -= instruction.num
		case "E":
			wpeast += instruction.num
		case "W":
			wpeast -= instruction.num
		}
	}

	return east, north
}

func main() {
	file, _ := os.Open("../input.txt")
	var direction string
	var num int
	instructions := make([]instruction, 0)
	for {
		_, err := fmt.Fscanf(file, "%1s%d", &direction, &num)
		if err == io.EOF {
			break
		}

		instructions = append(instructions, instruction{direction, num})
	}

	east, north := getFinalPosition(instructions)
	fmt.Println(math.Abs(float64(east)) + math.Abs(float64(north)))
}
