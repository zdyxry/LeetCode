package main

import "fmt"

func decode(encoded []int, first int) []int {
	res := []int{first}
	for i, d := range encoded {
		res = append(res, d^res[i])
	}
	return res
}

func main() {
	encode := []int{1, 2, 3}
	first := 1
	res := decode(encode, first)
	fmt.Println(res)
}
