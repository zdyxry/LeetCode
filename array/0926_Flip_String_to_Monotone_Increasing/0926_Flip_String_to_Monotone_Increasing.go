package main

import "fmt"

func minFlipsMonoIncr(S string) int {
	leftOnes := make([]int, len(S)+1)
	rightZeros := make([]int, len(S)+1)

	for index, ch := range S {
		leftOnes[index+1] = leftOnes[index]
		if ch == '1' {
			leftOnes[index+1]++
		}
	}

	for index := len(S) - 1; index >= 0; index-- {
		rightZeros[index] = rightZeros[index+1]
		if S[index] == '0' {
			rightZeros[index]++
		}
	}

	minFunc := func(a int, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}
	fmt.Println(leftOnes)
	fmt.Println(rightZeros)
	ans := minFunc(leftOnes[len(S)], rightZeros[0])
	for i := 0; i < len(S); i++ {
		turnCount := leftOnes[i+1] + rightZeros[i+1]
		ans = minFunc(ans, turnCount)
	}
	return ans
}

func main() {
	S := "00110"
	res := minFlipsMonoIncr(S)
	fmt.Println(res)
}
