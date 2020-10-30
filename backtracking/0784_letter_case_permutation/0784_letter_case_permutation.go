package main

import "fmt"

func letterCasePermutation(S string) []string {
	permutations := []string{S}
	for i, v := range S {
		if v >= 'a' && v <= 'z' {
			length := len(permutations)
			for ip := 0; ip < length; ip++ {
				b := []byte(permutations[ip])
				b[i] = byte(v + 'A' - 'a')
				permutations = append(permutations, string(b))
			}
		} else if v >= 'A' && v <= 'Z' {
			length := len(permutations)
			for ip := 0; ip < length; ip++ {
				b := []byte(permutations[ip])
				b[i] = byte(v + 'a' - 'A')
				permutations = append(permutations, string(b))
			}
		}
	}
	return permutations
}

func main() {
	S := "a1b2"
	res := letterCasePermutation(S)
	fmt.Println(res)
}
