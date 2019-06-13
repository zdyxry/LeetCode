package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make([]int, 26)
	for i, x := range s {
		m[x-'a']++
		m[t[i]-'a']--
	}

	for _, i := range m {
		if i != 0 {
			return false
		}
	}
	return true
}
func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))
}
