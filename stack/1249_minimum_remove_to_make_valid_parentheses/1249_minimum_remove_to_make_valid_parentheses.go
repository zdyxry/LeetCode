package main

import (
	"fmt"
	"strings"
)

func minRemoveToMakeValid(s string) string {
	stack := make([]int, 0, len(s))

	toDelete := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) == 0 {
				toDelete[i] = 1
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	for i := 0; i < len(stack); i++ {
		toDelete[stack[i]] = 1
	}

	var sb strings.Builder

	for i := 0; i < len(s); i++ {
		if toDelete[i] == 0 {
			sb.WriteByte(s[i])
		}
	}

	return sb.String()
}

func main() {
	s := "lee(t(c)o)de)"
	res := minRemoveToMakeValid(s)
	fmt.Println(res)
}
