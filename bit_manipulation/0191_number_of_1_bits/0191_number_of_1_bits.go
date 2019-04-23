package main

import "fmt"

func hammingBits(num uint32) int {
	var res uint32
	for num != 0 {
		res += (num & 1)
		num >>= 1
	}
	return int(res)
}

func hammingBits2(num uint32) int {
	var res int
	for num != 0 {
		res += 1
		num &= (num - 1)
	}
	return res
}

func main() {
	num := uint32(11)
	res := hammingBits(num)
	fmt.Println(res)
	res = hammingBits2(num)
	fmt.Println(res)
}
