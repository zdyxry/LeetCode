package main

import "fmt"

func minOperations(s string) int {
	var ans1, ans2 int
	for i := range s {
		if i%2 == 0 {
			if string(s[i]) != "0" {
				ans2++
			} else {
				ans1++
			}
		} else {
			if string(s[i]) != "0" {
				ans1++
			} else {
				ans2++
			}
		}
	}
	return min(ans1, ans2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := "0100"
	res := minOperations(s)
	fmt.Println(res)
}
