package main

import "os"
import "github.com/fredrik-jansson-se/go-advent-of-code-2018/day1"
import "github.com/fredrik-jansson-se/go-advent-of-code-2018/day2"
import "github.com/fredrik-jansson-se/go-advent-of-code-2018/day3"

func main() {
	var dmap = make(map[string]func())
	dmap["1"] = day1.Run
	dmap["2"] = day2.Run
	dmap["3"] = day3.Run
	dmap[os.Args[1]]()
}
