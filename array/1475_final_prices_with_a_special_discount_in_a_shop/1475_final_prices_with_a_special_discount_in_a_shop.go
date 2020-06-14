package main

import "fmt"

func finalPrices(prices []int) []int {
	ans := make([]int, len(prices))
	var stack []int
	for i, p := range prices {
		fmt.Println(i, p)
		for len(stack) > 0 {
			idx := stack[len(stack)-1]
			if p <= prices[idx] {
				ans[idx] = prices[idx] - p
				stack = stack[0 : len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}
	fmt.Println(stack)
	for _, idx := range stack {
		ans[idx] = prices[idx]
	}
	return ans
}

func main() {
	prices := []int{8, 4, 6, 2, 3}
	res := finalPrices(prices)
	fmt.Println(res)
}
