package main

import "fmt"

func maxScore(s string) int {
	var max, i, z int
	for _, v := range s {
		if v == '1' {
			i++
		}
	}

	end := len(s) - 1
	for _, v := range s[:end] {
		if v == '1' {
			i--
		} else {
			z++
		}
		m := i + z
		if max < m {
			max = m
		}
	}
	return max
}

func main() {
	s := "011101"
	res := maxScore(s)
	fmt.Println(res)
}
