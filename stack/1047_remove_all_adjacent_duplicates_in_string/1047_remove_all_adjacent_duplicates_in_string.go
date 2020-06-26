package main

import "fmt"

func removeDuplicates(S string) string {
	ans := make([]byte, 0, len(S))
	for _, v := range []byte(S) {
		if len(ans) > 0 && v == ans[len(ans)-1] {
			ans = ans[:len(ans)-1]
			continue
		}
		ans = append(ans, v)
	}

	return string(ans)
}

func main() {
	S := "abbaca"
	res := removeDuplicates(S)
	fmt.Println(res)
}
