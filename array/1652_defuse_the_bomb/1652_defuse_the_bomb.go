package main

import "fmt"

func decrypt(code []int, k int) []int {
	ans, l := []int{}, len(code)
	if k == 0 {
		return make([]int, l)
	}
	if k > 0 {
		for i := 0; i < l; i++ {
			tmp := 0
			for j := 1; j <= k; j++ {
				if i+j >= l {
					tmp += code[i+j-l]
				} else {
					tmp += code[i+j]
				}
			}
			ans = append(ans, tmp)
		}
	} else {
		for i := 0; i < l; i++ {
			tmp := 0
			for j := 1; j <= -k; j++ {
				if i-j <= -1 {
					tmp += code[i-j+l]
				} else {
					tmp += code[i-j]
				}
			}
			ans = append(ans, tmp)
		}
	}
	return ans
}

func main() {
	code := []int{5, 7, 1, 4}
	k := 3
	res := decrypt(code, k)
	fmt.Println(res)
}
