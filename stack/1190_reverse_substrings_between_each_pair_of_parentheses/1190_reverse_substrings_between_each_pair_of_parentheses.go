package main

import "fmt"

func reverseParentheses(s string) string {
	var res []byte
	for i := 0; i < len(s); i++ {
		if s[i] != ')' {
			res = append(res, s[i])
		} else {
			var temp []byte
			for res[len(res)-1] != '(' {
				temp = append(temp, res[len(res)-1])
				res = res[:len(res)-1]
			}

			res = res[:len(res)-1]
			for j := 0; j < len(temp); j++ {
				res = append(res, temp[j])
			}
		}
	}

	return string(res)
}

func main() {
	s := "(abcd)"
	res := reverseParentheses(s)
	fmt.Println(res)
}
