package main

import "fmt"

const mod = 1000000007

func numSub(s string) int {
	sum, count := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count++
		} else {
			sum = (sum + (count+1)*count/2) % mod
			count = 0
		}
	}
	if count != 0 {
		sum = (sum + (count+1)*count/2) % mod
	}
	return sum
}

func main() {
	s := "0110111"
	res := numSub(s)
	fmt.Println(res)
}
