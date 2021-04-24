package main 
import "fmt"
func minOperations(nums []int) int {
    cnt := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] <= nums[i-1]{
            cnt += nums[i-1] + 1 - nums[i]
            nums[i] = nums[i-1] + 1
        }
    }
    return cnt
    
}

func main() {
	nums := []int{1,1,1}
	res := minOperations(nums)
	fmt.Println(res)
}