package main 

import "fmt"

func subarraySum(a []int, k int) int {
	res, sum := 0, 0
	rec := make(map[int]int, len(a))
	rec[0] = 1

	for i := range a {
		sum += a[i]
		res += rec[sum-k]
		rec[sum]++
	}

	return res
}

func main() {
	a := []int{1,1,1}
	k := 2
	res := subarraySum(a, k)
	fmt.Println(res)
}