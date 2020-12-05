package main

import (
	"../lib"
	"fmt"
	"sort"
)

func main() {
	var lines []string

	readFunction := func(line string) {
		lines = append(lines, line)

	}

	lib.ReadFileByLine(5, readFunction)

	fmt.Println(Solve1(lines))
	fmt.Println(Solve2(lines))
}

func Solve1(iids []string) int {
	var sids []int
	for _, iid := range iids {
		_, _, sid := DecodeCard(iid)
		sids = append(sids, sid)
	}
	return MaxInt(sids)
}

func Solve2(iids []string) int {
	var sids []int
	for _, iid := range iids {
		_, _, sid := DecodeCard(iid)
		sids = append(sids, sid)
	}

	sort.Ints(sids)

	for i, n := range sids {
		if sids[i+1] == n+2 {
			return n + 1
		}
	}
	return -1
}

func DecodeCard(iid string) (row int, col int, seat int) {
	rowid := iid[0:7]
	colid := iid[7:]
	row = DecodeIidPart(rowid)
	col = DecodeIidPart(colid)
	seat = GetSeatId(row, col)
	return
}

func DecodeIidPart(iid string) int {
	n := 0
	d := len(iid) - 1

	for i, c := range iid {
		if c == 'B' || c == 'R' {
			n ^= 1 << (d - i)
		}
	}

	return n
}

func GetSeatId(row, col int) int {
	return row*8 + col
}

func MaxInt(nums []int) int {
	max := -1
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}
