package main 

import (
	"fmt"
)

func decodeAtIndex(S string, K int) string {
	lenSi, i := 0, 0

	for ; lenSi < K; i++ {
		char := S[i]
		if isDigit(char) {
			lenSi *= int(char - '0')

		} else {
			lenSi++
		}
	}

	for {
		i--
		char := S[i]
		if isDigit(char) {
			lenSi /= int(char - '0')
			K %= lenSi
		} else {
			if K == 0 || K == lenSi {
				return string(char)
			}
		}
		lenSi--
	}
}

func isDigit(b byte) bool {
	return b <= '9'
}

func main() {
	S := "leet2code3"
	K := 10
	res := decodeAtIndex(S, K)
	fmt.Println(res)
}