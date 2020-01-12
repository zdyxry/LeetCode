package main

import "fmt"

func numTrees(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 || n == 2 {
		return n
	}

	if n == 3 {
		return 5
	}

	res := 0
	for i := 1; i <= n/2; i++ {
		res += numTrees(i-1) * numTrees(n-i)
	}
	res *= 2
	if n%2 == 1 {
		temp := numTrees(n / 2)
		res += temp * temp
	}

	return res
}

func main() {
	n := 6
	res := numTrees(n)
	fmt.Println(res)
}
