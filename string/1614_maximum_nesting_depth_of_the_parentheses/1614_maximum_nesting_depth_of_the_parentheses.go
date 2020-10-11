package main

import "fmt"

func maxDepth(s string) int {
	max := 0
	var arr []string
	for i := 0; i < len(s); i++ {
		str := s[i : i+1]
		if str == "(" || str == ")" {
			if str == "(" {
				arr = append(arr, str)
				continue
			}
			if len(arr) > max {
				max = len(arr)
			}
			arr = arr[:len(arr)-1]
		}
	}
	return max
}

func main() {
	s := "(1+(2*3)+((8)/4))+1"
	res := maxDepth(s)
	fmt.Println(res)
}
