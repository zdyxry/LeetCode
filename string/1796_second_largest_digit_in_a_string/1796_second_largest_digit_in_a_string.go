package main

import "fmt"

func secondHighest(s string) int {
	var counts [10]int
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			counts[ch-'0'] = 1
		}
	}
	sign := false
	for i := 9; i >= 0; i-- {
		if counts[i] == 1 {
			if sign {
				return i
			}
			sign = true
		}
	}
	return -1
}

func main() {
	s := "dfa12321afd"
	res := secondHighest(s)
	fmt.Println(res)
}
