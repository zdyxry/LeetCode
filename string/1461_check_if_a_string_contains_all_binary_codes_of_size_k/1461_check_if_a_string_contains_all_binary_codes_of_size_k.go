package main

import (
	"fmt"
	"math"
)

func hasAllCodes(s string, k int) bool {
	m := map[string]bool{}
	for i := 0; i < len(s)-k+1; i++ {
		m[s[i:i+k]] = true
	}
	return len(m) == int(math.Pow(2, float64(k)))
}

func main() {
	s := "00110110"
	k := 2
	res := hasAllCodes(s, k)
	fmt.Println(res)
}
