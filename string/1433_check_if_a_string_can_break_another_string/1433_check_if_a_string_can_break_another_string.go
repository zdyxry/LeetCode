package main

import "fmt"

func checkIfCanBreak(s1 string, s2 string) bool {
	return helper(s1, s2) || helper(s2, s1)
}

func helper(s1 string, s2 string) bool {
	m2 := map[int]int{}
	for i := 0; i < len(s2); i++ {
		m2[int(s2[i]-'a')]++
	}
	for i := 0; i < len(s1); i++ {
		j := int(s1[i] - 'a')
		for ; j <= 26; j++ {
			if m2[j] > 0 {
				m2[j]--
				break
			}
		}
		if j > 26 {
			return false
		}

	}
	return true
}

func main() {
	s1 := "abc"
	s2 := "efg"
	res := checkIfCanBreak(s1, s2)
	fmt.Println(res)
}
