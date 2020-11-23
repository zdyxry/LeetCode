package main

import "fmt"

func getSmallestString(n int, k int) string {
	t := k - n
	ans := make([]byte, n, n)
	for i := n - 1; i >= 0; i-- {
		if t > 25 {
			ans[i] = 'z'
			t -= 25
		} else {
			ans[i] = byte('a' + t)
			t = 0
		}
	}

	return string(ans)
}

func main() {
	n := 3
	k := 27
	res := getSmallestString(n, k)
	fmt.Println(res)
}
