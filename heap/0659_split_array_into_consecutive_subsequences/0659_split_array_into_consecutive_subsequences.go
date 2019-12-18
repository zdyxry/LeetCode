package main 

import (
	"fmt"
)

func isPossible(nums []int) bool {
	counter := make(map[int]int)
	tail := make(map[int]int)
	for i := 0; i< len(nums); i++ {
		counter[nums[i]]++
	}

	for _, v := range nums {
		if counter[v] == 0{
			continue
		}

		if n := tail[v-1]; n> 0 {
			tail[v-1]--
			tail[v]++
			counter[v]--
			continue
		}
		if counter[v+1] > 0 && counter[v+2] > 0 {
			tail[v+2]++
			counter[v]--
			counter[v+1]--
			counter[v+2]--
		} else {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{1,2,3,3,4,5}
	res := isPossible(nums)
	fmt.Println(res)
}