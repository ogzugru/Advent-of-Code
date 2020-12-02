package main

import (
	. "../lib"
	"fmt"
	"strconv"
	"strings"
)

func lineSplitter(s string) (int, int, string, string) {
	// Split line on spaces
	splt := strings.Split(s, " ")

	// Get min and max occurances of character in password
	occSplit := strings.Split(splt[0], "-")
	i1, err := strconv.Atoi(occSplit[0])
	if err != nil {
		panic(err)
	}
	i2, err := strconv.Atoi(occSplit[1])
	if err != nil {
		panic(err)
	}

	// Get the character to count
	char := strings.Trim(splt[1], ":")
	return i1, i2, char, splt[2]
}
func main() {
	correctPW1 := 0
	correctPW2 := 0

	readFunction := func(line string) {

		correctPW1 += CheckPasswords_Part1(lineSplitter(line))
		correctPW2 += CheckPasswords_Part2(lineSplitter(line))

	}

	ReadFileByLine(2, readFunction)

	fmt.Println("Part 1: ", correctPW1)
	fmt.Println("Part 2: ", correctPW2)
}

func CheckPasswords_Part2(pos1 int, pos2 int, key string, pw string) int {
	correctPW := 0

	// Check if char is in place
	ok1 := string(pw[pos1-1]) == key
	ok2 := string(pw[pos2-1]) == key
	if (ok1 || ok2) && !(ok1 && ok2) {
		correctPW++
	}
	return correctPW
}

func CheckPasswords_Part1(minOcc int, maxOcc int, key string, pw string) int {
	correctPW := 0

	charOcc := strings.Count(pw, key)

	// Assert if password is correct
	if charOcc >= minOcc && charOcc <= maxOcc {
		correctPW++
	}

	return correctPW
}
