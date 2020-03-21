package main

import "fmt"

func kthGrammar(N int, K int) int {
	if K == 1 {
		return 0
	}

	if K%2 == 0 {
		source := kthGrammar(N-1, K/2)
		if source == 1 {
			return 0
		} else {
			return 1
		}
	} else {
		source := kthGrammar(N-1, (K+1)/2)
		if source == 1 {
			return 1
		} else {
			return 0
		}
	}
}

func main() {
	N := 1
	K := 1
	res := kthGrammar(N, K)
	fmt.Println(res)
}
