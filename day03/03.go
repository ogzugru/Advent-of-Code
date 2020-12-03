package main

import (
	"../lib"
	"fmt"
)

func countTrees(slopeX, slopeY int, strings []string) int {
	courseLength := len(strings)
	courseWidth := len(strings[0])
	treeCount := 0
	x := 0
	for y := slopeY; y < courseLength; y += slopeY {
		x = (x + slopeX) % courseWidth
		if string(strings[y][x]) == "#" {
			treeCount++
		}
	}

	return treeCount
}

func main() {
	var a []string

	readFunction := func(line string) {
		a = append(a, line)

	}

	lib.ReadFileByLine(3, readFunction)

	fmt.Println("Answer for Part 1: ", countTrees(3, 1, a))
	fmt.Println("Answer for Part 2: ", countTrees(1, 1, a)*countTrees(3, 1, a)*countTrees(5, 1, a)*countTrees(7, 1, a)*countTrees(1, 2, a))

}
