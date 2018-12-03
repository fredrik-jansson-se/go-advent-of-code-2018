package day1

import "fmt"
import "strconv"
import "github.com/fredrik-jansson-se/go-advent-of-code-2018/util"

func Run() {
	lines, err := util.ReadLines("day1.txt")
	if err != nil {
		panic(err)
	}

	deltas := make([]int, len(lines))
	for i, v := range lines {
		d, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		deltas[i] = d
	}

	run1(deltas)
	run2(deltas)
}

func run1(deltas []int) {
	sum := 0
	for _, v := range deltas {
		sum += v
	}
	fmt.Printf("1: %v\n", sum)
}

func run2(deltas []int) {
	sum := 0
	found := false
	sums := make(map[int]bool)
	sums[sum] = true
	for !found {
		for _, v := range deltas {
			sum += v
			if _, ok := sums[sum]; ok {
				found = true
				break
			}
			sums[sum] = true
		}
	}
	fmt.Printf("2: %v\n", sum)
}
