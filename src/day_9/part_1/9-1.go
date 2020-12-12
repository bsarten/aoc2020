package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func checkValid(q *list.List, num int) bool {
	for e := q.Front(); e != nil; e = e.Next() {
		for e2 := q.Front(); e2 != nil; e2 = e2.Next() {
			sum := e2.Value.(int) + e.Value.(int)
			if num == sum {
				return true
			}
		}
	}

	return false
}

func main() {
	b, _ := ioutil.ReadFile("../input.txt")
	index := 0
	q := list.New()
	for _, numStr := range strings.Split(string(b), "\n") {
		if numStr == "" {
			continue
		}

		num, _ := strconv.Atoi(numStr)
		if q.Len() >= 25 {
			if !checkValid(q, num) {
				fmt.Println(num)
				return
			}
			q.PushBack(num)
			q.Remove(q.Front())
		} else {
			q.PushBack(num)
		}
		index++
	}
}
