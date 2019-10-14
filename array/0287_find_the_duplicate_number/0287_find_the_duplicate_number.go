package main 

import "fmt"

func findDuplicate(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid, count := low+(high-low)>>1, 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count > mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func findDuplicate1(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		fmt.Println("slow is %d, fast is %d", slow, fast)
		slow, fast = nums[slow], nums[nums[fast]]
	}

	slow = 0
	for slow != fast {
		slow, fast = nums[slow], nums[fast]
	}
	return slow
}

func main() {
	nums := []int{1,3,4,2,2}
	res := findDuplicate(nums)
	res1 := findDuplicate1(nums)
	fmt.Println(res)
	fmt.Println(res1)
}