package main

// 1,2,3,5,4,6,
// 1,2,3,5,2,6

func canBeIncreasing(nums []int) bool {
	chance := true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			if chance == false {
				// no chance, return false
				return false
			}
			if i > 0 {
				// we have to remove nums[i+1] if nums[i+1] ≤ nums[i-1]
				if nums[i+1] <= nums[i-1] {
					nums[i+1] = nums[i]
				}
			}
			chance = false
		}
	}
	return true
}
