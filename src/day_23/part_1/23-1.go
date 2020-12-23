package main

import (
	"container/ring"
	"fmt"
)

func findInRing(r *ring.Ring, value int) *ring.Ring {
	currentLook := r.Next()
	for {
		if currentLook.Value.(int) == value {
			break
		}
		if currentLook == r {
			break
		}
		currentLook = currentLook.Next()
	}

	if currentLook == r {
		return nil
	}

	return currentLook
}

func printRing(r *ring.Ring) {
	for i := 0; i < r.Len()-1; i++ {
		fmt.Print(r.Value.(int))
		r = r.Next()
	}
	fmt.Println()
}

func main() {
	inputValues := []int{2, 8, 4, 5, 7, 3, 9, 6, 1}

	r := ring.New(len(inputValues))
	for _, inputValue := range inputValues {
		r.Value = inputValue
		r = r.Next()
	}

	for i := 0; i < 100; i++ {
		removedRing := r.Unlink(3)
		lookFor := r.Value.(int) - 1
		if lookFor < 1 {
			lookFor = 9
		}

		lookRing := r
		for lookRing == r {
			lookRing = findInRing(r, lookFor)
			if lookRing == nil {
				lookRing = r
				lookFor--
				if lookFor < 1 {
					lookFor = 9
				}
			}
		}
		lookRing.Link(removedRing)
		r = r.Next()
	}

	oneRing := findInRing(r, 1)
	oneRing = oneRing.Next()
	printRing(oneRing)
}
