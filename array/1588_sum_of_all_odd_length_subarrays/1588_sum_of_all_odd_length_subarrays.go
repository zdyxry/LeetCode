package main

import "fmt"

func sumOddLengthSubarrays(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	var count = 0
	var tmp = 0
	for i := 0; i < len(arr); i++ {
		tmp = arr[i]
		count = count + tmp
		for j := i + 1; j < len(arr); j++ {
			tmp = tmp + arr[j]
			if (j-i)%2 == 0 {
				count = count + tmp
			}
		}
	}
	return count
}

func main() {
	arr := []int{1, 4, 2, 5, 3}
	res := sumOddLengthSubarrays(arr)
	fmt.Println(res)
}
