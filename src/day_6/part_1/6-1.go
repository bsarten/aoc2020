package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	mapset "github.com/deckarep/golang-set"
)

func main() {
	b, _ := ioutil.ReadFile("../input.txt")
	sum := 0
	for _, group := range strings.Split(string(b), "\n\n") {
		yesAnswers := mapset.NewSet()
		for _, person := range strings.Split(group, "\n") {
			for _, answer := range person {
				yesAnswers.Add(answer)
			}
		}
		sum += yesAnswers.Cardinality()
	}
	fmt.Println(sum)
}
