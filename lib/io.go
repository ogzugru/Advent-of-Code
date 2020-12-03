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
