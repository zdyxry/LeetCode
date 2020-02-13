package main

import (
	"fmt"
)

func numOfSubarrays(arr []int, k int, threshold int) int {
	i := 0
	sum := 0
	count := 0
	length := len(arr)

	for i = 0; i < k; i++ {
		sum += arr[i]
	}
	for {
		if threshold*k <= sum {
			count++
		}

		if i >= length {
			break
		}
		sum = sum + arr[i] - arr[i-k]
		i++
	}
	return count
}

func main() {
	arr := []int{2, 2, 2, 2, 5, 5, 5, 8}
	k := 3
	threshold := 4
	res := numOfSubarrays(arr, k, threshold)
	fmt.Println(res)
}
