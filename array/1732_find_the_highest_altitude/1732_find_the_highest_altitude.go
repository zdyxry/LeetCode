package main

import "fmt"

func largestAltitude(gain []int) int {
	ans := 0
	ans = Max(ans, gain[0])
	for i := 1; i < len(gain); i++ {
		gain[i] += gain[i-1]
		ans = Max(ans, gain[i])
	}
	return ans
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	gain := []int{-5, 1, 5, 0, -7}
	res := largestAltitude(gain)
	fmt.Println(res)
}
