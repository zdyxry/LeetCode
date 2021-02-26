package main 

import "fmt"

func maxSatisfied(customers []int, grumpy []int, X int) int {
    res := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			res += customers[i]
			customers[i] = 0
		}
	}
	max_count := 0
	val := 0
	for i := 0; i < len(customers); i++ {
		if i < X{
			val += customers[i]
			continue
		}
		max_count = max(val, max_count)
		val += customers[i]
		val -= customers[i-X]
	}
	max_count = max(val, max_count)
	return res + max_count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	customers := []int{1,0,1,2,1,1,7,5}
	grupmy := []int{0,1,0,1,0,1,0,1}
	X := 3
	res := maxSatisfied(customers, grupmy, X)
	fmt.Println(res)
}