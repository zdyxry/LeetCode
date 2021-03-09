package main

import (
	"fmt"
	"strings"
)

func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}

func main() {
	s := "110"
	res := checkOnesSegment(s)
	fmt.Println(res)
}
