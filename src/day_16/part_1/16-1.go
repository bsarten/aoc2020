package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type fieldRange struct {
	low  int
	high int
}

type fieldDescription struct {
	name   string
	range1 fieldRange
	range2 fieldRange
}

func readFields(fieldsStr string) map[string]fieldDescription {
	fields := make(map[string]fieldDescription, 0)
	for _, field := range strings.Split(fieldsStr, "\n") {
		re := regexp.MustCompile(`^(.*): (\d+)-(\d+) or (\d+)-(\d+)$`)
		match := re.FindStringSubmatch(field)
		if len(match) == 6 {
			low1, _ := strconv.Atoi(match[2])
			high1, _ := strconv.Atoi(match[3])
			low2, _ := strconv.Atoi(match[4])
			high2, _ := strconv.Atoi(match[5])
			fields[match[1]] = fieldDescription{match[1], fieldRange{low1, high1}, fieldRange{low2, high2}}
		}
	}

	return fields
}

func fieldValid(number int, fields map[string]fieldDescription) bool {
	valid := false
	for _, v := range fields {
		if (v.range1.low <= number && v.range1.high >= number) || (v.range2.low <= number && v.range2.high >= number) {
			valid = true
			break
		}
	}
	return valid
}

func getErrorRate(ticketsStr string, fields map[string]fieldDescription) int {
	errorRate := 0
	for _, ticket := range strings.Split(ticketsStr, "\n")[1:] {
		for _, ticketField := range strings.Split(ticket, ",") {
			ticketNumber, _ := strconv.Atoi(ticketField)
			if !fieldValid(ticketNumber, fields) {
				errorRate += ticketNumber
			}
		}
	}
	return errorRate
}

func main() {
	b, _ := ioutil.ReadFile("../input.txt")
	sections := strings.Split(string(b), "\n\n")
	fields := readFields(sections[0])
	fmt.Println(getErrorRate(sections[2], fields))
}
