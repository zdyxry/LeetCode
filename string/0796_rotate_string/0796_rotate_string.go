package main

import "fmt"

func rotateString(A string, B string) bool {
	if A == B {
		return true
	}
	for i := 0; i < len(A); i++ {
		if B == string(A[i+1:])+string(A[0:i+1]) {
			return true
		}
	}
	return false
}

func main() {
	A := "abcde"
	B := "cdeab"
	res := rotateString(A, B)
	fmt.Println(res)
}
