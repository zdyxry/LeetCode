package main 

import "fmt"

func numSubarraysWithSum(A []int, S int) int {
	preSum, res := 0, 0
	count := make([]int, len(A)+1)
	count[0] = 1
	for _, n := range A {
		preSum += n
		if preSum >= S {
			res += count[preSum-S]
		}
		count[preSum]++
	}
	return res
}

func main() {
	A := []int{1,0,1,0,1}
	S := 2
	res := numSubarraysWithSum(A, S)
	fmt.Println(res)
}