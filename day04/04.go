package main

import (
	"../lib"
	"fmt"
	"regexp"
	"strings"
)

func isValidPassportPart1(line string) bool {

	var cre = map[string]string{"byr": "", "iyr": "", "eyr": "", "hgt": "", "hcl": "", "ecl": "", "pid": "", "cid": ""}

	var splitted []string
	splitted = strings.Split(line, " ")

	for i := 0; i < len(splitted); i++ {
		text := splitted[i][strings.Index(splitted[i], ":")+1:]
		switch sc := splitted[i][:strings.Index(splitted[i], ":")]; sc {
		case "byr":
			cre["byr"] = text
		case "iyr":
			cre["iyr"] = text
		case "eyr":
			cre["eyr"] = text
		case "hgt":
			cre["hgt"] = text
		case "hcl":
			cre["hcl"] = text
		case "ecl":
			cre["ecl"] = text
		case "pid":
			cre["pid"] = text
		case "cid":
			cre["cid"] = text
		default:
			fmt.Println("invalid suit %q", sc) // Joker?
		}
	}

	return cre["byr"] != "" && cre["iyr"] != "" && cre["eyr"] != "" && cre["hgt"] != "" && cre["hcl"] != "" && cre["ecl"] != "" && cre["pid"] != ""

}
func isValidPassport(line string) bool {

	var cre = map[string]string{"byr": "", "iyr": "", "eyr": "", "hgt": "", "hcl": "", "ecl": "", "pid": "", "cid": ""}

	var splitted []string
	splitted = strings.Split(line, " ")

	for i := 0; i < len(splitted); i++ {
		text := splitted[i][strings.Index(splitted[i], ":")+1:]
		switch sc := splitted[i][:strings.Index(splitted[i], ":")]; sc {
		case "byr":
			cre["byr"] = text
		case "iyr":
			cre["iyr"] = text
		case "eyr":
			cre["eyr"] = text
		case "hgt":
			cre["hgt"] = text
		case "hcl":
			cre["hcl"] = text
		case "ecl":
			cre["ecl"] = text
		case "pid":
			cre["pid"] = text
		case "cid":
			cre["cid"] = text
		default:
		}
	}

	return isValidPassportByByr(cre["byr"]) && isValidPassportByIyr(cre["iyr"]) && isValidPassportByEyr(cre["eyr"]) &&
		isValidPassportByHgt(cre["hgt"]) && isValidPassportByHcl(cre["hcl"]) && isValidPassportByEcl(cre["ecl"]) &&
		isValidPassportByPid(cre["pid"]) && isValidPassportByCid(cre["cid"])
}

func isValidPassportByByr(line string) bool {
	matched, _ := regexp.MatchString("^(19[2-9][0-9]|200[0-3])$", line)
	return matched
}

func isValidPassportByIyr(line string) bool {
	matched, _ := regexp.MatchString("^(201[0-9]|2020)$", line)
	return matched
}

func isValidPassportByEyr(line string) bool {
	matched, _ := regexp.MatchString("^(202[0-9]|2030)$", line)
	return matched
}

func isValidPassportByHgt(line string) bool {
	matched, _ := regexp.MatchString("^(1[5-8][0-9]cm|19[0-3]cm|59in|7[0-6]in|6[0-9]in)$", line)
	return matched
}
func isValidPassportByHcl(line string) bool {
	matched, _ := regexp.MatchString("^#[0-9a-f]{6}$", line)
	return matched
}

func isValidPassportByEcl(line string) bool {
	matched, _ := regexp.MatchString("^amb|blu|brn|gry|grn|hzl|oth$", line)
	return matched
}
func isValidPassportByPid(line string) bool {
	matched, _ := regexp.MatchString("^[0-9]{9}$", line)
	return matched
}
func isValidPassportByCid(line string) bool {

	return true
}

func main() {

	var k []string
	var temp string
	var count = 0

	readFunction := func(line string) {

		if "" == line {
			k = append(k, temp)
			temp = ""
			count++
		} else {
			if temp == "" {
				temp += line
			} else {
				temp += " "
				temp += line
			}

		}

	}

	lib.ReadFileByLine(4, readFunction)

	validatedCountPart1 := 0
	validatedCountPart2 := 0

	for _, v := range k {
		if isValidPassport(v) {
			validatedCountPart2++
		}
		if isValidPassportPart1(v) {
			validatedCountPart1++
		}

	}
	fmt.Println("Part 1: ", validatedCountPart1)
	fmt.Println("Part 2: ", validatedCountPart2)

}
