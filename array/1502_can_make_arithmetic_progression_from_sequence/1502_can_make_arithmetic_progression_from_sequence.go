package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	for i := 0; i < len(arr)-1; i++ {
		if i != 0 && arr[i+1]-arr[i] != arr[1]-arr[0] {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{1, 2, 4}
	res := canMakeArithmeticProgression(arr)
	fmt.Println(res)
}
