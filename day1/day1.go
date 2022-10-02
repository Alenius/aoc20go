package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/alenius/aoc20go/lib"
)

func readTestData() map[int]int {
	testdata := map[int]int{
		1721: 1721,
		979:  979,
		366:  366,
		299:  299,
		675:  675,
		1456: 1456,
	}
	return testdata
}

func main() {
	linemap := map[int]int{}
	linefunc := func(line string) {
		asint, _ := strconv.Atoi(line)
		linemap[asint] = asint
	}
	if err := lib.ReadFileByLine("day1", linefunc); err != nil {
		fmt.Println(err)
	}
	if os.Getenv("TEST") != "" {
		linemap = readTestData()
	}

	find := func(m map[int]int, val int) int {
		for k := range m {
			v := val - k
			_, ok := m[v]
			if ok {
				return k * v
			}
		}
		return -1
	}

	fmt.Printf("pt1: %v \n", find(linemap, 2020))

	pt2 := func() int {
		for k := range linemap {
			v := find(linemap, 2020-k)
			if v != -1 {
				return k * v
			}
		}
		return -1
	}
	fmt.Printf("pt2 %v \n", pt2())
}
