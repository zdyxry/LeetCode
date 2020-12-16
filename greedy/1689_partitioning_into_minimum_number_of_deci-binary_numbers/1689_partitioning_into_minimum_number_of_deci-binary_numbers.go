package main

import "fmt"

func minPartitions(n string) int {
	ans := byte(0)
	for i := 0; i < len(n); i++ {
		if n[i] > ans {
			ans = n[i]
		}
	}
	return int(ans - '0')
}

func main() {
	n := "32"
	res := minPartitions(n)
	fmt.Println(res)
}
