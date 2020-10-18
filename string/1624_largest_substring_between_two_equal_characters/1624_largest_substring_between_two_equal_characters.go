package main

import "fmt"

func maxLengthBetweenEqualCharacters(s string) int {
	m := make(map[rune]int)
	max := -1
	for i, ch := range s {
		if pos, ok := m[ch]; ok {
			if i-pos-1 > max {
				max = i - pos - 1
			}
		} else {
			m[ch] = i
		}
	}
	return max
}

func main() {
	s := "aa"
	res := maxLengthBetweenEqualCharacters(s)
	fmt.Println(res)
}
