package main

import "fmt"

func validMountainArray(A []int) bool {
	length := len(A)
	if length < 3 {
		return false
	}
	count := 0 // 转折点 本题山脉应当是 左上坡 右下坡 且 只有一个转折点
	pre := 0   // 前一个点的趋势 1为上坡 -1为下坡
	for i := 0; i < length-1; i++ {
		if A[i] == A[i+1] {
			return false
		}
		if A[i] > A[i+1] && i == 0 {
			return false
		}
		if A[i] < A[i+1] { // 上坡
			if pre == -1 {
				return false
			}
			pre = 1
		}
		if A[i] > A[i+1] { // 下坡
			if pre == 1 {
				count++
			}
			pre = -1
		}
	}
	return count == 1
}

func main() {
	A := []int{1, 2}
	res := validMountainArray(A)
	fmt.Println(res)
}
