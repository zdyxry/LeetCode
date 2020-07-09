package main

import (
	"fmt"
	"math"
	"sort"
)

func longestStrChain(words []string) int {
	if len(words) < 2 {
		return len(words)
	}

	var dp = make([]int, len(words))
	var res int
	dp[0] = 1
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) < len(words[j]) {
			return true
		} else {
			return false
		}
	})

	for i := 1; i < len(words); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if isChain(words[j], words[i]) {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
		res = int(math.Max(float64(res), float64(dp[i])))
	}
	return res
}

func isChain(a, b string) bool {
	if len(b)-len(a) == 1 {
		var i, j int
		for i < len(a) && j < len(b) {
			if a[i] == b[j] {
				i++
			}
			j++
		}
		if i == len(a) {
			return true
		}
	}
	return false
}

func main() {
	words := []string{"a", "b", "ba", "bca", "bda", "bdca"}
	res := longestStrChain(words)
	fmt.Println(res)
}
