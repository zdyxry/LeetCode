package main

import "fmt"

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	if len(s1) != len(s2) {
		return false
	}
	num := 0
	dif := []int{}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			num += 1
			dif = append(dif, i)
		}
	}
	if num != 2 {
		return false
	}
	if s1[dif[0]] == s2[dif[1]] && s1[dif[1]] == s2[dif[0]] {
		return true
	}
	return false
}

func main() {
	s1 := "bank"
	s2 := "kanb"
	res := areAlmostEqual(s1, s2)
	fmt.Println(res)
}
