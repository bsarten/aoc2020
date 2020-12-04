package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func validateByr(fields *map[string]string) bool {
	_, valid := (*fields)["byr"]
	return valid
}

func validateIyr(fields *map[string]string) bool {
	_, valid := (*fields)["iyr"]
	return valid
}

func validateEyr(fields *map[string]string) bool {
	_, valid := (*fields)["eyr"]
	return valid
}

func validateHgt(fields *map[string]string) bool {
	_, valid := (*fields)["hgt"]
	return valid
}

func validateHcl(fields *map[string]string) bool {
	_, valid := (*fields)["hcl"]
	return valid
}

func validateEcl(fields *map[string]string) bool {
	_, valid := (*fields)["ecl"]
	return valid
}

func validatePid(fields *map[string]string) bool {
	_, valid := (*fields)["pid"]
	return valid
}

func validatePassport(fields *map[string]string) bool {
	if validateByr(fields) && validateIyr(fields) && validateEyr(fields) && validateHgt(fields) && validateHcl(fields) && validateEcl(fields) && validatePid(fields) {
		return true
	}
	return false
}

func main() {
	file, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	validPassports := 0
	fields := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if validatePassport(&fields) {
				validPassports++
			}
			fields = make(map[string]string)
			continue
		}

		for _, field := range strings.Split(line, " ") {
			nv := strings.Split(field, ":")
			name := nv[0]
			value := nv[1]
			fields[name] = value
		}
	}
	fmt.Println(validPassports)
}
