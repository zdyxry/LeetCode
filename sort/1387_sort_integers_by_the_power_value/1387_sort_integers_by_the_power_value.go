package main

import (
	"fmt"
	"sort"
)

func getKth(lo int, hi int, k int) int {
	power := make(map[int]int)
	arr := make([]int, hi-lo+1)
	for i := lo; i <= hi; i++ {
		arr[i-lo] = i
		getPower(i, power)
	}
	sort.Slice(arr, func(i, j int) bool {
		if power[arr[i]] != power[arr[j]] {
			return power[arr[i]] < power[arr[j]]
		}
		return arr[i] < arr[j]
	})
	return arr[k-1]
}

func getPower(x int, power map[int]int) int {
	ans, tmp := 0, x
	for x != 1 {
		if x&1 == 1 {
			x = 3*x + 1
		} else {
			x /= 2
		}
		ans++
		if val, ok := power[x]; ok {
			ans += val
			break
		}
	}
	power[tmp] = ans
	return ans
}

func main() {
	lo := 12
	hi := 15
	k := 2
	res := getKth(lo, hi, k)
	fmt.Println(res)
}
