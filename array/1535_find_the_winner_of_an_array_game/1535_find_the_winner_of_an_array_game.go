package main

import "fmt"

func getWinner(arr []int, k int) int {
	size := len(arr)
	ans, cnt := arr[0], 0
	for i := 1; i < size; i++ {
		if arr[i] > ans {
			ans, cnt = arr[i], 1
		} else {
			cnt++
		}
		if cnt == k {
			return ans
		}
	}
	return ans
}

func main() {
	arr := []int{2, 1, 3, 5, 4, 6, 7}
	k := 2
	res := getWinner(arr, k)
	fmt.Println(res)
}
