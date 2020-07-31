package main

import "fmt"

const mod = int(1e9 + 7)

func numOfSubarrays(arr []int) int {
	odd, even := 0, 0
	res := 0
	for _, v := range arr {
		if v%2 != 0 {
			odd, even = even+1, odd
		} else {
			odd, even = odd, even+1
		}
		res += odd
		res %= mod
	}
	return res
}

func main() {
	arr := []int{1, 3, 5}
	res := numOfSubarrays(arr)
	fmt.Println(res)
}
