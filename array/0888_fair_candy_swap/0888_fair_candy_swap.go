package main

import "fmt"

func fairCandySwap(A []int, B []int) []int {
	var diff int
	for _, num := range A {
		diff += num
	}
	for _, num := range B {
		diff -= num
	}
	diff /= 2
	data := make(map[int]struct{}, len(A))
	for _, num := range A {
		data[num] = struct{}{}
	}
	for _, num := range B {
		if _, ok := data[num+diff]; ok {
			return []int{num + diff, num}
		}
	}
	return nil
}

func main() {
	A := []int{1, 1}
	B := []int{2, 2}
	res := fairCandySwap(A, B)
	fmt.Println(res)
}
