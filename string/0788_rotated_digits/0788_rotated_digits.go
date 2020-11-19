package main

import "fmt"

func good(n int, flag bool) bool {
	if n == 0 {
		return flag
	}
	d := n % 10
	switch d {
	case 3, 4, 7:
		return false
	case 0, 1, 8:
		return good(n/10, flag)
	}
	//2,5,6,9
	return good(n/10, true)
}

func rotatedDigits(N int) int {
	flag := false
	ans := 0
	for i := 1; i <= N; i++ {
		if good(i, flag) {
			ans++
		}
	}
	return ans
}

func main() {
	N := 10
	res := rotatedDigits(N)
	fmt.Println(res)
}
