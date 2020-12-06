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

	allAnswers := mapset.NewSet()
	for i := 'a'; i <= 'z'; i++ {
		allAnswers.Add(i)
	}

	for _, group := range strings.Split(string(b), "\n\n") {
		groupAnswers := allAnswers.Clone()
		for _, person := range strings.Split(group, "\n") {
			personAnswers := mapset.NewSet()
			for _, answer := range person {
				personAnswers.Add(answer)
			}
			groupAnswers = groupAnswers.Intersect(personAnswers)
		}
		sum += groupAnswers.Cardinality()
	}
	fmt.Println(sum)
}
