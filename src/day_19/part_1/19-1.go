package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func readRule(rules map[string]string, input string) {
	colonSplit := strings.Split(input, ": ")
	ruleNumber := colonSplit[0]
	ruleStr := colonSplit[1]
	rules[ruleNumber] = ruleStr
}

func expandRuleNumbers(rules map[string]string, rule string) string {
	expandedStr := ""
	numberSplit := strings.Split(rule, " ")
	for _, nextRuleNumber := range numberSplit {
		expandedStr += expandRule(rules, nextRuleNumber)
	}

	return expandedStr
}

func expandRule(rules map[string]string, ruleNumber string) string {
	rule := rules[ruleNumber]
	if rule[0] == '"' {
		return rule[1 : len(rule)-1]
	}

	pipeSplit := strings.Split(rule, " | ")
	left := pipeSplit[0]
	var right string
	if len(pipeSplit) == 2 {
		right = pipeSplit[1]
	}

	left = expandRuleNumbers(rules, left)
	ruleStr := "(" + left
	if right != "" {
		right = expandRuleNumbers(rules, right)
		ruleStr += "|" + right
	}
	ruleStr += ")"

	rules[ruleNumber] = "\"" + ruleStr + "\""
	return ruleStr
}

func messageMatchesRule(rules map[string]string, ruleNumber string, message string) bool {
	rule := rules[ruleNumber]
	r := regexp.MustCompile("^" + rule[1:len(rule)-1] + "$")
	return r.MatchString(message)
}

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sections := strings.Split(string(b), "\n\n")
	ruleLines := sections[0]
	rulesInput := strings.Split(ruleLines, "\n")
	rules := make(map[string]string, len(rulesInput))
	for _, ruleInput := range rulesInput {
		readRule(rules, ruleInput)
	}

	expandRule(rules, "0")

	count := 0
	messageLines := strings.Split(sections[1], "\n")
	for _, message := range messageLines {
		if messageMatchesRule(rules, "0", message) {
			count++
		}
	}

	fmt.Println(count)
}
