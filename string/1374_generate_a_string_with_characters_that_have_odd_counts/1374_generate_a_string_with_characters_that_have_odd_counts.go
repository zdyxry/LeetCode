package main

import (
	"fmt"
	"strings"
)

func generateTheString(n int) string {
	var rst string
	for i := 0; n != 0 && i < 26; i++ {
		if n%2 == 1 {
			rst += strings.Repeat(string('a'+i), n)
			n = 0
		} else {
			rst += strings.Repeat(string('a'+i), n-1)
			n = 1
		}
	}
	return rst
}

func main() {
	n := 4
	res := generateTheString(n)
	fmt.Println(res)
}
