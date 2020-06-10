package main

import "fmt"

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	for n := 1; ; n++ {
		sub := n*(n+1)/2 - target
		if sub >= 0 && sub&1 == 0 {
			return n
		}
	}
}

func main() {
	target := 1
	res := reachNumber(target)
	fmt.Println(res)
}
