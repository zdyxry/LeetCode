package main

import "fmt"

func diStringMatch(S string) []int {
	var (
		l, r = 0, len(S)
		ans  = make([]int, 0, len(S)+1)
	)
	for _, s := range S {
		switch s {
		case 'I':
			ans = append(ans, l)
			l++
		default:
			ans = append(ans, r)
			r--
		}
	}
	ans = append(ans, l)

	return ans
}

func main() {
	S := "IDID"
	res := diStringMatch(S)
	fmt.Println(res)
}
