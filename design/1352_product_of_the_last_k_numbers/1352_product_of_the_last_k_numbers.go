package main 

import "fmt"

type ProductOfNumbers struct {
	nums []int
}

func Constructor() ProductOfNumbers {
	pn := ProductOfNumbers{
		nums: []int{},
	}
	return pn
}

func (this *ProductOfNumbers) Add(num int) {
	this.nums = append(this.nums, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	res := 1
	len := len(this.nums)
	for i :=0; i < k; i++ {
		res *= this.nums[len-1-i]
		if res == 0 {
			return 0
		}
	}
	return res
}
