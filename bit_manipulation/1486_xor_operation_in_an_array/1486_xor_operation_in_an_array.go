package main

import "fmt"

func xorOperation(n int, start int) (xor int) {
	for i := 0; i < n; i++ {
		xor = xor ^ (start + 2*i)
	}
	return
}

func main() {
	n := 5
	start := 0
	res := xorOperation(n, start)
	fmt.Println(res)
}
