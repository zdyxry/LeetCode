package main

import "fmt"

func truncateSentence(s string, k int) string {
	var blankCount int
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			break
		}
		if s[i] == ' ' {
			blankCount++
		}
		if blankCount == k {
			return s[0:i]
		}
	}
	return s
}

func main() {
	s := "Hello how are you Contestant"
	k := 4
	res := truncateSentence(s, k)
	fmt.Println(res)
}
