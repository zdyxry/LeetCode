package main 

import "fmt"

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T), len(T))
	stack := make([]int, 1, len(T))

	for i := 1; i < len(T); i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack) - 1]] {
			idx := stack[len(stack)-1]
			res[idx] = i - idx
			stack = stack[:len(stack) - 1]

		}
		stack = append(stack, i)
	}

	return res
}

func main() {
	T := []int{73,74,75,71,69,72,76,73}
	res := dailyTemperatures(T)
	fmt.Println(res)
}