package main

import "fmt"

func totalHammingDistance(nums []int) int {
	ans, numsLen := 0, len(nums)
	for i := 0; i < 32; i++ {
		haveOne := 0
		for _, num := range nums {
			haveOne += (num >> i) & 1
		}
		ans += haveOne * (numsLen - haveOne)
	}
	return ans
}

func main() {
	nums := []int{4, 14, 2}
	res := totalHammingDistance(nums)
	fmt.Println(res)
}
