package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	m := map[int64]struct{}{}
	readFunction := func(line string) {
		n, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		m[n] = struct{}{}
	}
	ReadFileByLine(1, readFunction)

	p1, _ := seek(m, 2020, -1)
	WriteSolutionPart1("%d", p1)

	for k, _ := range m {
		p2, ok := seek(m, 2020-k, k)
		if ok {
			WriteSolutionPart2("%d", k*p2)
			break
		}
	}
}
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

func WriteSolutionPart1(format string, values ...interface{}) {
	answer := "Part 1: " + format + "\n"
	fmt.Printf(answer, values...)
}

func WriteSolutionPart2(format string, values ...interface{}) {
	answer := "Part 2: " + format + "\n"
	fmt.Printf(answer, values...)
}

func seek(m map[int64]struct{}, n, ignore int64) (int64, bool) {
	for k, _ := range m {
		if k == ignore {
			continue
		}
		v := n - k
		if _, ok := m[v]; ok {
			return k * v, true
		}
	}
	return -1, false
}
