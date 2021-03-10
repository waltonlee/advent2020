package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input []string
var passports []map[string]string

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	var count int
	passport := map[string]string{}

	for _, v := range input {
		if v == "" && len(passport) > 0 {
			passports = append(passports, passport)
			passport = make(map[string]string)
		} else {
			for _, fields := range strings.Split(v, " ") {
				pair := strings.Split(fields, ":")
				passport[pair[0]] = pair[1]
			}
		}
	}
	if len(passport) > 0 {
		passports = append(passports, passport)
	}

	for _, passport := range passports {
		var byrok, iyrok, eyrok, hgtok, hclok, eclok, pidok bool
		byr, byrb := passport["byr"]
		byri, err := strconv.Atoi(byr)
		if byrb && err == nil && byri >= 1920 && byri <= 2002 {
			byrok = true
		}
		iyr, iyrb := passport["iyr"]
		iyri, err := strconv.Atoi(iyr)
		if iyrb && err == nil && iyri >= 2010 && iyri <= 2020 {
			iyrok = true
		}
		eyr, eyrb := passport["eyr"]
		eyri, err := strconv.Atoi(eyr)
		if eyrb && err == nil && eyri >= 2020 && eyri <= 2030 {
			eyrok = true
		}
		hgt, hgtb := passport["hgt"]
		var hgtn int
		var hgtt string
		_, err = fmt.Sscanf(hgt, "%d%s", &hgtn, &hgtt)
		if hgtb && err == nil && ((hgtt == "cm" && hgtn >= 150 && hgtn <= 193) || (hgtt == "in" && hgtn >= 59 && hgtn <= 76)) {
			hgtok = true
		}
		hcl, hclb := passport["hcl"]
		if hclb && len(hcl) == 7 && string(hcl[0]) == "#" {
			hclok = true
		}
		ecl, eclb := passport["ecl"]
		if eclb && (ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth") {
			eclok = true
		}
		pid, pidb := passport["pid"]
		if _, err := strconv.Atoi(pid); pidb && err == nil && len(pid) == 9 {
			pidok = true
		}

		if byrok && iyrok && eyrok && hgtok && hclok && eclok && pidok {
			count++
		}
	}

	fmt.Println(count)
}
