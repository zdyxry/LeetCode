package main

import "fmt"

func maximumPopulation(logs [][]int) int {
	var year [101]int
	for _, l := range logs {
		year[l[0]-1950]++
		year[l[1]-1950]--
	}
	cnt := 0
	mcnt := 0
	ret := 0
	for y, c := range year {
		cnt += c
		if cnt > mcnt {
			mcnt = cnt
			ret = y + 1950
		}
	}
	return ret
}

func main() {
	logs := [][]int{{1993, 1999}, {2000, 2010}}
	res := maximumPopulation(logs)
	fmt.Println(res)
}
