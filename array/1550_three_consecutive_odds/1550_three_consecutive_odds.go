package main

import "fmt"

func threeConsecutiveOdds(arr []int) bool {
	cnt := 0
	for _, n := range arr {
		if n&1 == 1 {
			cnt++
		} else {
			cnt = 0
		}
		if cnt == 3 {
			return true
		}
	}
	return false

}

func main() {
	arr := []int{1, 2, 34, 3, 4, 5, 7, 23, 12}
	res := threeConsecutiveOdds(arr)
	fmt.Println(res)
}
