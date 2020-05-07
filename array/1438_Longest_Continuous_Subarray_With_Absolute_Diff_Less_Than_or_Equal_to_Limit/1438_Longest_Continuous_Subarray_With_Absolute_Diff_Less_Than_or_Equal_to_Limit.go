package main 

import "fmt"

func longestSubarray(nums []int, limit int) int {
    minD := []int{} // queue for mins -> increasing order of mins from start to end
    maxD := []int{} // queue for maxs -> decresing order of maxs from start to end
    start := 0
    end := 0
    for end < len(nums) {
        current := nums[end]
        for len(minD) > 0 && last(minD) > current {
            minD = trim(minD)
        }
        for len(maxD) > 0 && last(maxD) < current {
            maxD = trim(maxD)
        }
        minD = append(minD, current)
		maxD = append(maxD, current)
		fmt.Println(minD, maxD)
		
        
        // Note: At this point, maxD has max, and minD has min numbers for start to end
        if len(maxD) > 0 && len(minD) > 0 && (maxD[0]-minD[0] > limit) {
            if maxD[0]==nums[start] {
                maxD=maxD[1:]
            }
            if minD[0]==nums[start] {
                minD=minD[1:]
            }
            start++
        }
        end++
    }
    return (end-start)
}

func last(a []int) int {
    return a[len(a)-1]
}

func trim(a []int) []int {
    return a[:len(a)-1]
}

func main() {
	nums := []int{8,2,4,7}
	limit := 4
	res := longestSubarray(nums, limit)
	fmt.Println(res)
}