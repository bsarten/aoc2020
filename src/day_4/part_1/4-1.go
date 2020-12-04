package main

import (
	"fmt"
	"io/ioutil"
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
	b, _ := ioutil.ReadFile("../input.txt")
	validPassports := 0
	for _, passport := range strings.Split(string(b), "\n\n") {
		passport := strings.ReplaceAll(passport, "\n", " ")
		fields := make(map[string]string)
		for _, field := range strings.Split(passport, " ") {
			nv := strings.Split(field, ":")
			if len(nv) == 2 {
				name := nv[0]
				value := nv[1]
				fields[name] = value
			}
		}

		if validatePassport(&fields) {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}
