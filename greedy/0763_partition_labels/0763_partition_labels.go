package main

import "fmt"

func partitionLabels(S string) []int {
	// Use a 26 length slice to
	// store all leters' frequency.
	lastIndex := make([]int, 26)
	for i, c := range S {
		lastIndex[int(c-'a')] = i
	}
	ans := make([]int, 0)
	start := 0
	end := 0
	for i := 0; i < len(S); i++ {
		end = max(lastIndex[int(S[i]-'a')], end)
		if i == end {
			ans = append(ans, end-start+1)
			start = end + 1
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	S := "ababcbacadefegdehijhklij"
	res := partitionLabels(S)
	fmt.Println(res)
}
