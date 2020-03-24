package main

import "math"

func sumFourDivisors(nums []int) int {
	summary := 0
	for _, v := range nums {
		summary += helper(v)
	}
	return summary
}

func helper(num int) int {
	if num <= 5 {
		return 0
	}
	root := int(math.Sqrt(float64(num)))
	if root*root == num {
		return 0
	}
	divisorCount := 2
	summary := num + 1
	for i := 2; i <= root; i++ {
		if num%i == 0 {
			divisorCount += 2
			summary += i
			summary += num / i
			if divisorCount > 4 {
				return 0
			}
		}
	}
	if divisorCount == 4 {
		return summary
	}
	return 0
}
