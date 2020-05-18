package main

import "fmt"

func simplifiedFractions(n int) []string {
	ret := []string{}
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(i, j) == 1 {
				ret = append(ret, fmt.Sprintf("%d/%d", j, i))
			}
		}
	}
	return ret
}

func gcd(x, y int) int {
	for x%y != 0 {
		x, y = y, x%y
	}
	return y
}

func main() {
	n := 3
	res := simplifiedFractions(n)
	fmt.Println(res)
}
