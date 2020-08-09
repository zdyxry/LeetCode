package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	prev, curr, cnt, inc := 0, 0, 0, 0
	for _, curr = range arr {
		inc = curr - prev - 1
		if cnt+inc >= k {
			break
		}
		prev = curr
		cnt += inc
	}
	inc = k - cnt
	return prev + inc
}

func main() {
	arr := []int{1, 2, 3, 4}
	res := findKthPositive(arr, 2)
	fmt.Println(res)
}
