package main 

import (
	"fmt"
)

func removeKdigits(num string, k int) string {
	digits := len(num) - k 
	stack := make([]byte, len(num))
	top := 0

	for i := range num {
		for top > 0 && stack[top-1] > num[i] && k > 0 {
			top--
			k--
		}
		stack[top] = num[i]
		top++
	}

	i := 0
	for i < digits && stack[i] == '0' {
		i++
	}
	if i == digits {
		return "0"
	}
	return string(stack[i:digits])
}



func main() {
	num := "1432219"
	k := 3
	res := removeKdigits(num, k)
	fmt.Println(res)
}