package main

import "fmt"

func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}

	nums := make([]int, 300)
	res := 0

	nums[0] = 0
	nums[1] = 1
	for i := 1; i <= n; i++ {
		nums[i*2] = nums[i]
		nums[i*2+1] = nums[i] + nums[i+1]
		if nums[i] > res {
			res = nums[i]
		}
	}

	return res
}

func main() {
	n := 7
	res := getMaximumGenerated(n)
	fmt.Println(res)
}
