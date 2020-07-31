package main

import "fmt"

func minFlips(target string) int {
	var count int
	target = "0" + target
	flag := target[len(target)-1]
	for i := len(target) - 1; i >= 0; i-- {
		if target[i] != flag {
			count++
			flag = target[i]
		}
	}

	return count
}

func main() {
	target := "10111"
	res := minFlips(target)
	fmt.Println(res)
}
