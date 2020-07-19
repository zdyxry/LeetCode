package main

import "fmt"

func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	}
	return 0.5
}

func main() {
	n := 1
	res := nthPersonGetsNthSeat(n)
	fmt.Println(res)
}
