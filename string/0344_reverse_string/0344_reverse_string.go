package main

import (
	"fmt"
)

func reverseString(s []byte) {
	start := 0
	end := len(s) - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
func main() {
	s := []byte("hello")
	fmt.Println(string(s))
	reverseString(s)
	fmt.Println(string(s))
}
