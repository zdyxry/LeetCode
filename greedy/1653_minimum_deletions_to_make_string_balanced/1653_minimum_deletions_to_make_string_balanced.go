package main

import "fmt"

func minimumDeletions(s string) int {
	n := len(s)
	a := make([]int, n, n)
	b := make([]int, n, n)

	if s[0] == 'a' {
		a[0] = 1
	}
	if s[n-1] == 'b' {
		b[n-1] = 1
	}

	for i := 1; i < n; i++ {
		a[i] = a[i-1]
		if s[i] == 'a' {
			a[i]++
		}
	}

	for j := n - 2; j >= 0; j-- {
		b[j] = b[j+1]
		if s[j] == 'b' {
			b[j]++
		}
	}
	// fmt.Println(a, b)
	res := 0

	for i := 0; i < n; i++ {
		x := a[i] + b[i]
		if x > res {
			res = x
		}
	}

	return n - res
}

func main() {
	s := "aababbab"
	res := minimumDeletions(s)
	fmt.Println(res)
}
