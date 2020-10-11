package main

import "fmt"

func minimumOneBitOperations(n int) int {
	var count, turn int

	for i := 31; i >= 0; i-- {
		if n&(1<<i) > 0 {
			if turn&1 == 0 {
				count += (1 << (i + 1)) - 1
			} else {
				count -= (1 << (i + 1)) - 1
			}
			turn++
		}
	}

	return count
}

func main() {
	n := 5
	res := minimumOneBitOperations(n)
	fmt.Println(res)
}
