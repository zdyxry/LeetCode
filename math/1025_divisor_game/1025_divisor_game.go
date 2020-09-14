package main

import "fmt"

func divisorGame(N int) bool {
	return N%2 == 0
}

func main() {
	N := 2
	res := divisorGame(N)
	fmt.Println(res)
}
