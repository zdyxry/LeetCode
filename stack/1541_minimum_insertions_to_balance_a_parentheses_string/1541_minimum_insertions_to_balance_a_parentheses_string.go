package main

import "fmt"

func minInsertions(s string) int {
	r := 0

	stack := 0

	i := 0
	for i < len(s) {
		if s[i] == '(' {
			stack++
			i++
		} else { // ')'
			i++
			if i == len(s) {
				r++
				if stack == 0 {
					r++
				} else {
					stack--
				}
			} else {
				if s[i] != ')' {
					r++
				} else {
					i++
				}

				if stack == 0 {
					r++
				} else {
					stack--
				}

			}
		}
	}

	return r + stack*2
}

func main() {
	s := "(()))"
	res := minInsertions(s)
	fmt.Println(res)
}
