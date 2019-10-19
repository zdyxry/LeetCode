package main 

import "fmt"

func singleNumber(nums []int) []int {
	var xor int
	for _, num := range nums {
		xor ^= num
	}
	
	lowest := xor & -xor
	fmt.Println(xor, lowest)
	var a,b int
	for _, num := range nums {
		if num&lowest == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a,b}
}

func main() {
	nums := []int{1,2,1,3,2,5}
	res := singleNumber(nums)
	fmt.Println(res)
}