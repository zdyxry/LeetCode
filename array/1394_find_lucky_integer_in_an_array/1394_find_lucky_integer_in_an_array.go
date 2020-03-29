package main

import "fmt"

func findLucky(arr []int) int {
	freqCounter := make(map[int]int)

	for _, elem := range arr {
		freqCounter[elem]++
	}

	result := -1
	for key, value := range freqCounter {
		if key == value && result < value {
			result = value
		}
	}
	return result
}

func main() {
	arr := []int{2, 2, 3, 4}
	res := findLucky(arr)
	fmt.Println(res)
}
