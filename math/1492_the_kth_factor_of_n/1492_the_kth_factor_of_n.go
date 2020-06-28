package main

import "fmt"

func kthFactor(n int, k int) int {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			k--
			if k == 0 {
				return i
			}
		}
	}
	return -1
}

func main() {
	n := 12
	k := 3
	res := kthFactor(n, k)
	fmt.Println(res)
}
