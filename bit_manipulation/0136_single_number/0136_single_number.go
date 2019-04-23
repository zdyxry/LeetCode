package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res ^= n
	}
	return res

}

func main() {
	nums := []int{1, 2, 2, 3, 1}
	res := singleNumber(nums)
	fmt.Println(res)
}
