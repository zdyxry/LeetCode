package main

import "fmt"

func sumOfUnique(nums []int) int {
	counter := make(map[int]int)
	for _, x := range nums {
		if _, ok := counter[x]; !ok {
			counter[x] = 1
		} else {
			counter[x]++
		}
	}
	res := 0
	for k, v := range counter {
		if v == 1 {
			res += k
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 2}
	res := sumOfUnique(nums)
	fmt.Println(res)
}
