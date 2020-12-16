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
	canBe  map[int]bool
}

func readFields(fieldsStr string) map[string]fieldDescription {
	fields := make(map[string]fieldDescription, 0)
	ticketFields := strings.Split(fieldsStr, "\n")
	for _, field := range strings.Split(fieldsStr, "\n") {
		re := regexp.MustCompile(`^(.*): (\d+)-(\d+) or (\d+)-(\d+)$`)
		match := re.FindStringSubmatch(field)
		if len(match) == 6 {
			low1, _ := strconv.Atoi(match[2])
			high1, _ := strconv.Atoi(match[3])
			low2, _ := strconv.Atoi(match[4])
			high2, _ := strconv.Atoi(match[5])
			canBe := make(map[int]bool, 0)
			for i := 0; i < len(ticketFields); i++ {
				canBe[i] = true
			}
			fields[match[1]] = fieldDescription{match[1], fieldRange{low1, high1}, fieldRange{low2, high2}, canBe}
		}
	}

	return fields
}

func numValidForField(number int, field fieldDescription) bool {
	if (field.range1.low <= number && field.range1.high >= number) || (field.range2.low <= number && field.range2.high >= number) {
		return true
	}
	return false
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

func isTicketValid(ticketStr string, fields map[string]fieldDescription) bool {
	valid := true
	for _, ticketField := range strings.Split(ticketStr, ",") {
		ticketNumber, _ := strconv.Atoi(ticketField)
		if !fieldValid(ticketNumber, fields) {
			valid = false
			break
		}
	}
	return valid
}

func canBeOnlyOne(field fieldDescription) (bool, int) {
	canBe := -1
	for k, v := range field.canBe {
		if v {
			if canBe != -1 {
				return false, 0
			}
			canBe = k
		}
	}
	return true, canBe
}

func lockOrder(fields map[string]fieldDescription) {
	locked := make(map[string]bool, len(fields))
	for k := range fields {
		locked[k] = false
	}

	numlocked := 0
	for numlocked != len(fields) {
		numlocked = 0
		for fieldIdx, field := range fields {
			if locked[fieldIdx] {
				numlocked++
				continue
			}
			if onlyOne, idx := canBeOnlyOne(field); onlyOne {
				numlocked++
				locked[fieldIdx] = true
				for k, field := range fields {
					if k != fieldIdx {
						field.canBe[idx] = false
					}
				}
			}
		}
	}
}

func getFieldOrder(ticketsStr string, fields map[string]fieldDescription) []string {
	for _, ticket := range strings.Split(ticketsStr, "\n")[1:] {
		if !isTicketValid(ticket, fields) {
			continue
		}

		for idx, ticketField := range strings.Split(ticket, ",") {
			ticketNumber, _ := strconv.Atoi(ticketField)
			for _, field := range fields {
				if !numValidForField(ticketNumber, field) {
					field.canBe[idx] = false
				}
			}
		}
	}

	lockOrder(fields)

	fieldOrder := make([]string, len(fields))
	for _, field := range fields {
		for k, v := range field.canBe {
			if v {
				fieldOrder[k] = field.name
			}
		}
	}
	return fieldOrder
}

func multiplyDepartures(ticketStr string, fieldOrder []string) int {
	product := 1
	ticketLines := strings.Split(ticketStr, "\n")
	ticketFields := strings.Split(ticketLines[1], ",")
	for i := 0; i < len(fieldOrder); i++ {
		fieldname := fieldOrder[i]
		if strings.Contains(fieldname, "departure") {
			number, _ := strconv.Atoi(ticketFields[i])
			product *= number
		}
	}
	return product
}

func main() {
	b, _ := ioutil.ReadFile("../input.txt")
	sections := strings.Split(string(b), "\n\n")
	fields := readFields(sections[0])
	fieldOrder := getFieldOrder(sections[2], fields)
	fmt.Println(multiplyDepartures(sections[1], fieldOrder))
}
