package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func validateByr(fields *map[string]string) bool {
	valid := false
	year, _ := strconv.Atoi((*fields)["byr"])
	if year >= 1920 && year <= 2002 {
		valid = true
	}
	return valid
}

func validateIyr(fields *map[string]string) bool {
	valid := false
	year, _ := strconv.Atoi((*fields)["iyr"])
	if year >= 2010 && year <= 2020 {
		valid = true
	}
	return valid
}

func validateEyr(fields *map[string]string) bool {
	valid := false
	year, _ := strconv.Atoi((*fields)["eyr"])
	if year >= 2020 && year <= 2030 {
		valid = true
	}
	return valid
}

func validateHgt(fields *map[string]string) bool {
	valid := false
	re := regexp.MustCompile(`^([0-9]+)(cm|in)$`)
	match := re.FindStringSubmatch((*fields)["hgt"])
	if len(match) == 3 {
		num, _ := strconv.Atoi(match[1])
		incm := match[2]
		if incm == "cm" {
			if num >= 150 && num <= 193 {
				valid = true
			}
		} else if incm == "in" {
			if num >= 59 && num <= 76 {
				valid = true
			}
		}
	}
	return valid
}

func validateHcl(fields *map[string]string) bool {
	valid := false
	r := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	if r.MatchString((*fields)["hcl"]) {
		valid = true
	}
	return valid
}

func validateEcl(fields *map[string]string) bool {
	valid := false
	re := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	if re.MatchString((*fields)["ecl"]) {
		valid = true
	}
	return valid
}

func validatePid(fields *map[string]string) bool {
	valid := false
	r := regexp.MustCompile(`^[0-9]{9}$`)
	if r.MatchString((*fields)["pid"]) {
		valid = true
	}
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
