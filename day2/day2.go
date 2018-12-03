package day2

import "fmt"
import "github.com/fredrik-jansson-se/go-advent-of-code-2018/util"

func Run() {
	lines, err := util.ReadLines("day2.txt")
	if err != nil {
		panic(err)
	}

	run1(lines)
}

type CharCount map[rune]uint

func run1(lines []string) {

	twoCnt := 0
	threeCnt := 0

	for _, l := range lines {
		var cnt = make(CharCount)
		for _, c := range l {
			if sum, ok := cnt[c]; ok {
				cnt[c] = sum + 1
			} else {
				cnt[c] = 1
			}
		}
		hasTwo := false
		hasThree := false
		for _, v := range cnt {
			if v == 2 {
				hasTwo = true
			} else if v == 3 {
				hasThree = true
			}
		}
		if hasTwo {
			twoCnt += 1
		}
		if hasThree {
			threeCnt += 1
		}
	}

	fmt.Printf("day2-1: %v\n", twoCnt*threeCnt)
}
