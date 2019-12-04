package main 

import (
	"fmt"
)

func find132pattern(a []int) bool {
	ak := -1 << 31
	ajStack := make([]int, 0, len(a))

	for i := len(a) - 1; 0 <= i; i-- {
		if a[i] < ak {
			return true
		}

		for len(ajStack) > 0 &&
			ajStack[len(ajStack)-1] < a[i] {
				ak = ajStack[len(ajStack)-1]
				ajStack = ajStack[:len(ajStack)-1]
		}
	ajStack = append(ajStack, a[i])
	}
	return false
}

func main() {
	a := []int{1,2,3,4}
	res := find132pattern(a)
	fmt.Println(res)
}