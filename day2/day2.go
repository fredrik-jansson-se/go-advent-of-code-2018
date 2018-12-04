package day2

import "fmt"
import "github.com/fredrik-jansson-se/go-advent-of-code-2018/util"

func Run() {
	lines, err := util.ReadLines("day2.txt")
	if err != nil {
		panic(err)
	}

	run1(lines)
	run2(lines)
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

func run2(lines []string) {
	for i1, l1 := range lines {
		for _, l2 := range lines[i1+1:] {
			diff := 0
			for i := 0; i < len(l1) && diff < 2; i++ {
				if l1[i] != l2[i] {
					diff += 1
				}
			}
			if diff < 2 {
				fmt.Printf("%v - %v\n", l1, l2)
				for i := 0; i < len(l1) && diff < 2; i++ {
					if l1[i] == l2[i] {
						fmt.Printf("%c", l1[i])
					}
				}

				fmt.Printf("\n")
				return
			}
		}
	}
}
