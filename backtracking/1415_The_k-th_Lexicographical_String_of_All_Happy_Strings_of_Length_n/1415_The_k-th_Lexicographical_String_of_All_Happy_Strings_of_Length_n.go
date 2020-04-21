package main

import "fmt"

func getHappyString(n int, k int) string {
	res := ""
	num := 0
	var dfs func(i int, buf []byte)
	dfs = func(i int, buf []byte) {
		if i == n {
			num++
			if num == k {
				res = string(buf)
			}
			return
		}
		for j := byte('a'); j < 'd'; j++ {
			if i == 0 || j != buf[i-1] {
				buf[i] = j
				dfs(i+1, buf)
			}
		}
	}
	dfs(0, make([]byte, n))
	return res
}

func main() {
	n := 1
	k := 3
	res := getHappyString(n, k)
	fmt.Println(res)
}
