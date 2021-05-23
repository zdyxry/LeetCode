package main 

import "fmt"

func checkZeroOnes(s string) bool {
	cnt0 := 0
	cnt1 := 0
	result0 := 0
	result1 := 0
	for i:=0; i < len(s); i++ {
		if s[i] == '0' {
			cnt0++
			cnt1 = 0
		} else {
			cnt1++
			cnt0 = 0
		}
		result0 = max(result0, cnt0)
		result1 = max(result1, cnt1)
	}
	return result1 > result0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "1101"
	res := checkZeroOnes(s)
	fmt.Println(res)
}