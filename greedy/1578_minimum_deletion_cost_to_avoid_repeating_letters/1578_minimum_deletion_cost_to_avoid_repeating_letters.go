package main

import "fmt"

func minCost(s string, cost []int) int {
	sum := 0
	for i := range cost {
		sum += cost[i]
	}
	res := 0
	for i := 0; i < len(s); {
		m := cost[i]
		if i < len(s)-1 && s[i] == s[i+1] {
			j := i
			for j < len(s) && s[j] == s[i] {
				m = max(m, cost[j])
				j++
			}
			i = j
		} else {
			i++
		}
		res += m
	}
	return sum - res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	s := "abaac"
	cost := []int{1, 2, 3, 4, 5}
	res := minCost(s, cost)
	fmt.Println(res)
}
