package main

import (
	"container/ring"
	"fmt"
)

func isInRing(r *ring.Ring, value int) bool {
	currentLook := r
	i := r.Len()
	for {
		if currentLook.Value.(int) == value {
			break
		}
		currentLook = currentLook.Next()
		i--
		if i == 0 {
			return false
		}
	}

	return true
}

func printRing(r *ring.Ring) {
	for i := 0; i < r.Len(); i++ {
		fmt.Print(r.Value.(int))
		r = r.Next()
	}
	fmt.Println()
}

func main() {
	valueMap := make(map[int]*ring.Ring, 0)
	inputValues := []int{2, 8, 4, 5, 7, 3, 9, 6, 1}

	r := ring.New(1000000)
	for _, inputValue := range inputValues {
		r.Value = inputValue
		valueMap[inputValue] = r
		r = r.Next()
	}

	for i := 10; i <= 1000000; i++ {
		r.Value = i
		valueMap[i] = r
		r = r.Next()
	}

	for i := 0; i < 10000000; i++ {
		removedRing := r.Unlink(3)

		lookFor := r.Value.(int) - 1
		if lookFor < 1 {
			lookFor = 1000000
		}

		for isInRing(removedRing, lookFor) {
			lookFor--
			if lookFor < 1 {
				lookFor = 1000000
			}
		}

		destinationRing := valueMap[lookFor]
		destinationRing.Link(removedRing)
		r = r.Next()
	}

	oneRing := valueMap[1]
	oneRing = oneRing.Next()
	result := oneRing.Value.(int)
	oneRing = oneRing.Next()
	result *= oneRing.Value.(int)
	fmt.Println(result)
}
