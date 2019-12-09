package main 

import (
	"fmt"
)

func sumSubarrayMins(A []int) int {
	stack, dp, res, mod := []int{}, make([]int, len(A)+1), 0, 1000000007
	stack = append(stack, -1)

	for i :=0; i < len(A); i++ {
		for stack[len(stack)-1] != -1 && A[i] <= A[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		dp[i+1] = (dp[stack[len(stack)-1]+1] + (i-stack[len(stack)-1])*A[i]) % mod
		stack = append(stack, i)
		res += dp[i+1]
		res %= mod
	}
	return res
}

func main() {
	A := []int{3,1,2,4}
	res := sumSubarrayMins(A)
	fmt.Println(res)
}