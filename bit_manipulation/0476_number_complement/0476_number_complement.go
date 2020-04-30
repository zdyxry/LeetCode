package main

import "fmt"

func findComplement(num int) int {
	temp := 1
	for ; temp <= num; temp *= 2 {
	}
	return (temp - 1) ^ num
}

func main() {
	num := 5
	res := findComplement(num)
	fmt.Println(res)
}
