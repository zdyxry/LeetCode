package main

import "fmt"

func prefixesDivBy5(A []int) []bool {
	num, result := 0, make([]bool, len(A))
	for k, v := range A {
		num = num<<1 + v
		if num%5 == 0 {
			result[k] = true
		}
		num %= 10

	}
	return result
}

func main() {
	A := []int{0, 1, 1}
	res := prefixesDivBy5(A)
	fmt.Println(res)
}
