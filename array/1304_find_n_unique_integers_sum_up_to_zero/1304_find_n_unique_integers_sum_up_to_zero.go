package main

import "fmt"

func sumZero(n int) []int {
	res := make([]int, n)
	for num, i := n/2, 0; i < n/2; i, num = i+1, num-1 {
		res[i] = -num
		res[n/2+i] = num
	}

	return res
}

func main() {
	n := 5
	res := sumZero(n)
	fmt.Println(res)
}
