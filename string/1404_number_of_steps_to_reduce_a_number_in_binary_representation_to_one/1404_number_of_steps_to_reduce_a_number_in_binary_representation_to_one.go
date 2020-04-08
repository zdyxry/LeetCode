package main

import "fmt"

func reverse(A []byte) {
	for i := 0; i < len(A)/2; i++ {
		A[i], A[len(A)-1-i] = A[len(A)-1-i], A[i]
	}
}

func numSteps(s string) int {
	if len(s) == 1 {
		return 0
	}
	bs := []byte(s)
	reverse(bs)
	step, carry, last := 0, 0, bs[len(bs)-1]
	bs = bs[:len(bs)-1]
	for _, v := range bs {
		if carry == 0 {
			if v == '1' {
				step += 2
				carry = 1
			} else {
				step += 1
			}
		} else {
			if v == '0' {
				step += 2

			} else {
				step += 1
				carry = 1
			}
		}
	}
	if last == '1' {
		step += carry
	}
	return step
}

func main() {
	s := "1101"
	res := numSteps(s)
	fmt.Println(res)
}
