package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type containsRule struct {
	num int
	bag string
}

func canContain(rules *map[string][]containsRule, bag string, target string) bool {
	bagRules := (*rules)[bag]
	for _, rule := range bagRules {
		if rule.bag == target {
			return true
		} else if rule.num != 0 {
			if canContain(rules, rule.bag, target) {
				return true
			}
		}
	}
	return false
}

func buildRulesMap(filename string) *map[string][]containsRule {
	b, _ := ioutil.ReadFile(filename)
	rulesMap := make(map[string][]containsRule)
	for _, rule := range strings.Split(string(b), "\n") {
		re := regexp.MustCompile(`^(.*) bags contain (.*).$`)
		match := re.FindStringSubmatch(string(rule))
		if len(match) == 3 {
			bagType := match[1]
			containsList := match[2]
			for _, contains := range strings.Split(containsList, ", ") {
				words := strings.Split(contains, " ")
				number, _ := strconv.Atoi(words[0])
				bagname := words[1] + " " + words[2]
				rulesMap[bagType] = append(rulesMap[bagType], containsRule{number, bagname})
			}
		}
	}

	return &rulesMap
}

func main() {

	rulesMap := buildRulesMap("../input.txt")
	num := 0
	for bag := range *rulesMap {
		if canContain(rulesMap, bag, "shiny gold") {
			num++
		}
	}
	fmt.Println(num)
}
