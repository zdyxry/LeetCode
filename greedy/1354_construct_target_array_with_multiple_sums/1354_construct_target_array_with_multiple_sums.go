package main

import (
	"fmt"
)

func isPossible(target []int) bool {
	if len(target) == 0 {
		return true
	}
	for j := 0; j < len(target); j++ {
		sum := 0
		maxIdx := 0
		maxNum := target[0]
		for i := 0; i < len(target); i++ {
			if target[i] > maxNum {
				maxIdx = i
				maxNum = target[i]
			}
			sum += target[i]
		}
		if maxNum == 1 {
			return true
		}
		if sum-maxNum >= maxNum || (sum-maxNum) == 0 {
			return false
		}
		target[maxIdx] = maxNum % (sum - maxNum)
	}
	return true
}

func main() {
	target := []int{1, 1000000}
	res := isPossible(target)
	fmt.Println(res)
}
