package main 

import "fmt"

func titleToNumber(s string) int {
	res := 0

	for i:=0; i< len(s); i ++ {
		temp := int(s[i] - 'A' + 1)
		res = res*26 + temp
	}

	return res
}

func main() {
	s := "A"
	res := titleToNumber(s)
	fmt.Println(res)
}