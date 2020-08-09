package main

import "fmt"

func makeGood(s string) string {
	ans := []byte{}
	for i := 0; i < len(s); i++ {
		ans = append(ans, s[i])
		size := len(ans)
		if size >= 2 {
			if mequal(ans[size-1], ans[size-2]) {
				ans = ans[:size-2]
			}
		}
	}
	return string(ans)
}

func mequal(a, b byte) bool {
	return (a ^ ' ') == b
}

func main() {
	s := "leEeetcode"
	res := makeGood(s)
	fmt.Println(res)
}
