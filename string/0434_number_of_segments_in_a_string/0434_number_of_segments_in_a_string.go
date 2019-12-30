package main

import (
	"fmt"
	"strings"
)

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}

	strs := strings.Fields(s)
	return len(strs)
}

func main() {
	s := "Hello, my name is John"
	res := countSegments(s)
	fmt.Println(res)
}
