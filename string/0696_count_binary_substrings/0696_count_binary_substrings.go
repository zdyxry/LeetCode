package main

import (
	"fmt"
)

func countBinarySubstrings(s string) int {
	count, countZero, countOne := 0, 0, 0
	prev := rune(s[0])

	for _, r := range s {
		if prev == r {
			if r == '0' {
				countZero++
			} else {
				countOne++
			}
		} else {
			count += min(countZero, countOne)
			if r == '0' {
				countZero = 1
			} else {
				countOne = 1
			}
		}
		prev = r
	}

	return count + min(countZero, countOne)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := "001100"
	fmt.Println(countBinarySubstrings(s))
}
