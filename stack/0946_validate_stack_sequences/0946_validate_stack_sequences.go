package main 

import (
	"fmt"
)

func validateStackSequences(pushed []int, poped []int) bool {
	size := len(pushed)

	stack := make([]int, size)
	top := -1

	for in, out := 0, 0; in < size; in++ {
		if pushed[in] != poped[out] {
			top++
			stack[top] = pushed[in]
		} else {
			out++
			for top >= 0 && stack[top] == poped[out] {
				top--
				out++
			}
		}
	}
	return top == -1
}

func validateStackSequences2(pushed, poped []int) bool {
	i := 0
	s := make([]int, 0, len(pushed))
	for _, v :=range pushed{
		s = append(s, v)
		for len(s) != 0 && i < len(poped) && s[len(s)-1] == poped[i] {
			s = s[:len(s)-1]
			i++
		}
	}
	return i == len(poped)
}

func main() {
	pushed := []int{1,2,3,4,5}
	poped := []int{4,5,3,2,1}
	// res := validateStackSequences(pushed, poped)
	res := validateStackSequences2(pushed, poped)
	fmt.Println(res)
}