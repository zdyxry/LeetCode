package main 

import (
	"fmt"
)

func majorityElement(nums []int) []int {
	count1, count2, candidate1, candidate2 := 0, 0, 0, 1
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 <= 0 {
			candidate1, count1 = num, 1
		} else if count2 <= 0 {
			candidate2, count2 = num ,1
		} else {
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2{
			count2++
		}
	}

	length := len(nums)
	if count1 > length/3 && count2 > length/3 {
		return []int{candidate1, candidate2}
	}

	if count1 > length/3 {
		return []int{candidate1}

	}

	if count2 > length / 3 {
		return []int{candidate2}
	}

	return []int{}
}

func main() {
	nums := []int{1,1,1,3,3,2,2,2}
	res := majorityElement(nums)
	fmt.Println(res)
}