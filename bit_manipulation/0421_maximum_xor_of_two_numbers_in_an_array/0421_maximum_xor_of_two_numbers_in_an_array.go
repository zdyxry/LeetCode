package main 

import "fmt"

func findMaximumXOR(nums []int) int {
	var max, mask int

	for i := 31; i >= 0; i-- {
		mask |= 1 << uint(i)
		
		nMap := make(map[int]bool)
		for _, num := range nums {
			nMap[num&mask] = true
		}

		tmp := max | (1 << uint32(i))
		for key := range nMap {
			if nMap[tmp^key] {
				max = tmp
				break
			}
		}
	}

	return max
}

func main() {
	nums := []int{3,10,5,25,2,8}
	res := findMaximumXOR(nums)
	fmt.Println(res)
}