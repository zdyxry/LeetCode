package main

import "fmt"

func minSteps(s, t string) int {
	var commonChar = 0
	var charA = map[uint8]int{}

	for i := 0; i < len(s); i++ {
		charA[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		charA[t[i]]--
	}

	for _, v := range charA {
		commonChar += abs(v)
	}
	return commonChar / 2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	s := "bab"
	t := "aba"
	res := minSteps(s, t)
	fmt.Println(res)
}
