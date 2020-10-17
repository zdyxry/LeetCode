package main

import "fmt"

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	ans := []int{0, 1}
	for i := 1; i < n; i++ {
		var next []int
		b := 1 << uint(i)
		for j := len(ans) - 1; j >= 0; j-- {
			next = append(next, ans[j]|b)
		}
		fmt.Println(next)
		ans = append(ans, next...)
	}
	return ans
}

func main() {
	n := 2
	res := grayCode(n)
	fmt.Println(res)
}
