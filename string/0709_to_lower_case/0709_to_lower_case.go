package main

import "fmt"

func toLowerCase(str string) string {
	runes := []rune(str)
	diff := 'a' - 'A'
	for i := 0; i < len(str); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += diff
		}
	}
	return string(runes)
}

func main() {
	str := "Hello"
	res := toLowerCase(str)
	fmt.Println(res)
}
