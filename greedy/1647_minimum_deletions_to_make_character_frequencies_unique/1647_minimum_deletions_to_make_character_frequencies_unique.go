package main

import (
	"fmt"
	"sort"
)

// count
func minDeletions(s string) int {

	counts := make([]int, 26)
	for i := range s {
		counts[s[i]-'a']++
	}

	sort.Ints(counts)

	res := 0
	for i := len(counts) - 2; i > -1; i-- {
		if counts[i] > 0 && counts[i] >= counts[i+1] {
			target := max(0, counts[i+1]-1)
			res += (counts[i] - target)
			counts[i] = target
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "aab"
	res := minDeletions(s)
	fmt.Println(res)
}
