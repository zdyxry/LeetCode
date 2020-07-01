package main

import "fmt"

func canArrange(arr []int, k int) bool {
	mod := make([]int, k)
	n := len(arr)
	for i := 0; i < n; i++ {
		mod[(arr[i]%k+k)%k]++
	}
	if mod[0]%2 > 0 {
		return false
	}
	for i := 1; i <= k/2; i++ {
		if 2*i == k {
			if mod[i]%2 > 0 {
				return false
			}
		} else {
			if mod[i] != mod[k-i] {
				return false
			}
		}
	}
	return true
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}
	k := 5
	res := canArrange(arr, k)
	fmt.Println(res)
}
