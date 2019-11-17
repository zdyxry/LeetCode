package main 

import "fmt"

func checkValidString(s string) bool {
	l, r := 0, 0
	n := len(s)
	for i:= 0;i<n;i++ {
		if s[i] == ')' {
			l--
		} else {
			l++
		}

		j := n-i-1 
		if s[j] == '(' {
			r--
		} else {
			r++
		}

		if l < 0 || r < 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "(*)"
	res := checkValidString(s)
	fmt.Println(res)
}