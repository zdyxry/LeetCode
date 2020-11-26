package main

import (
	"fmt"
	"strings"
)

func checkRecord(s string) bool {
	if strings.Count(s, "A") <= 1 && strings.Count(s, "LLL") == 0 {
		return true
	}
	return false
}

func main() {
	s := "PPALLP"
	res := checkRecord(s)
	fmt.Println(res)
}
