package main

import "fmt"

func countTriplets(arr []int) int {
	n := len(arr)
	prefix := make([]int, n+1, n+1)
	prefix[0] = 0
	for i := 1; i <= n; i++ {
		prefix[i] = arr[i-1] ^ prefix[i-1]
	}

	total := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if prefix[i] == prefix[j] {
				total += j - i - 1
			}
		}
	}
	return total
}

func main() {
	arr := []int{2, 3, 1, 6, 7}
	res := countTriplets(arr)
	fmt.Println(res)
}
