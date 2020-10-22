package main

import "fmt"

func shortestToChar(S string, C byte) []int {
	var l, r int
	ans := []int{}
	for i := 0; i < len(S); i++ {
		l = i
		r = i
		mindist := 0
		for {
			if S[l] == C || S[r] == C {
				ans = append(ans, mindist)
				break
			}
			if l-1 >= 0 {
				l--
			}
			if r+1 < len(S) {
				r++
			}
			mindist++
		}
	}
	return ans
}

func main() {
	S := "loveleetcode"
	C := byte('e')
	res := shortestToChar(S, C)
	fmt.Println(res)
}
