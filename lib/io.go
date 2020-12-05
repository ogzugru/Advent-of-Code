package lib

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// The file io.go is created for handling the lined file inputs.

func ReadFile(day int) string {
	filename := fmt.Sprintf("Day%02d/%02d.in", day, day)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func ReadFileByLine(day int, fn func(string)) {
	filename := fmt.Sprintf("Day%02d/%02d.in", day, day)
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fn(scanner.Text())
	}
}
func Remove(s []string, in string) []string {
	for i, v := range s {
		if v == in {
			s = append(s[:i], s[i+1:]...)
			break
		}
	}
	return s
}
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func FindMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
