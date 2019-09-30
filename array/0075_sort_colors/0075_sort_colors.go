package main 

import "fmt"

func sortColors(nums []int) {
	i, j, k := 0, 0, len(nums) -1
	for j <= k {
		switch nums[j] {
		case 0:
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		case 1:
			j++
		case 2:
			nums[j], nums[k] = nums[k], nums[j]
			k--
		}
	}
}

func main() {
	nums := []int{2,0,2,1,1,0}
	sortColors(nums)
	fmt.Println(nums)
}