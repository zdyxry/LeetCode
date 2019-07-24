package main

import "fmt"

func backspaceCompare(S, T string) bool {
	i := len(S)
	j := len(T)

	for i >= 0 || j >= 0 {
		fmt.Println("i and j:", i, j)
		i = nextIndex(&S, i)
		j = nextIndex(&T, j)
		fmt.Println(i, j, S[i], T[j])
		if i >= 0 && j >= 0 && S[i] != T[j] {
			return false
		}
	}

	return i == j
}

func nextIndex(s *string, i int) int {
	i--
	count := 0
	for i >= 0 && ((*s)[i] == '#' || count > 0) {
		if (*s)[i] == '#' {
			count++
		} else {
			count--
		}
		i--
	}
	return i
}

func main() {
	S := "abd#c"
	T := "abbd##c"
	result := backspaceCompare(S, T)
	fmt.Println(result)
}
