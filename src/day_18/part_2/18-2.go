package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Top() interface{} {
	return s.top.value
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

func readToken(expression string) (string, string) {
	if expression == "" {
		return "", ""
	}
	switch expression[0] {
	case ')':
		return ")", strings.TrimLeft(expression[1:], " ")
	case '(':
		return "(", strings.TrimLeft(expression[1:], " ")
	case '*':
		return "*", strings.TrimLeft(expression[1:], " ")
	case '+':
		return "+", strings.TrimLeft(expression[1:], " ")
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		re := regexp.MustCompile(`^(\d+)\s*(.*)$`)
		match := re.FindStringSubmatch(expression)
		return match[1], match[2]
	}

	return "", expression
}

func performOp(op string, operands *Stack) {
	num1, _ := strconv.Atoi(operands.Pop().(string))
	num2, _ := strconv.Atoi(operands.Pop().(string))
	switch op {
	case "*":
		operands.Push(fmt.Sprintf("%d", num1*num2))
	case "+":
		operands.Push(fmt.Sprintf("%d", num1+num2))
	}
}

func evaluate(expression string) int {

	var operators Stack
	var operands Stack

	var token string
	token, expression = readToken(expression)
	for token != "" {
		if token == "(" {
			operators.Push(token)
		} else if token == ")" {
			for !operators.Empty() {
				operator, _ := operators.Top().(string)
				if operator == "(" {
					break
				}
				performOp(operator, &operands)
				operators.Pop()
			}
			operators.Pop()
		} else if token != "*" && token != "+" {
			operands.Push(token)
		} else {
			for !operators.Empty() {
				operator, _ := operators.Top().(string)
				if operator == "(" || operator != "+" {
					break
				}
				performOp(operator, &operands)
				operators.Pop()
			}
			operators.Push(token)
		}
		token, expression = readToken(expression)
	}

	for !operators.Empty() {
		operator, _ := operators.Pop().(string)
		performOp(operator, &operands)
	}

	ret, _ := strconv.Atoi(operands.Pop().(string))

	return ret
}

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		sum += evaluate(line)
	}

	fmt.Println(sum)
}
