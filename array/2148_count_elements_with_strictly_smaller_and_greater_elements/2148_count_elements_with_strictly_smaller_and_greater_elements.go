package main 

func countElements(nums []int) int {
    sort.Ints(nums)
    
    result := 0
    for idx := range nums {
        if nums[0] < nums[idx]  && nums[idx] < nums[len(nums)-1] {
            result += 1
        }
    }
    return result
    
}