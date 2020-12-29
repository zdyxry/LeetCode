package main

import (
	"fmt"
	"strings"
)

func maximumBinaryString(binary string) string {

	firstZero := strings.IndexByte(binary, '0')
	if firstZero == -1 || firstZero == len(binary)-1 {
		return binary
	}
	res := []byte(binary)
	ones := strings.Count(binary[firstZero:], "1")
	for i := range res {
		res[i] = '1'
	}
	res[len(res)-1-ones] = '0'
	return string(res)
}

func main() {
	binary := "000110"
	res := maximumBinaryString(binary)
	fmt.Println(res)
}
