package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	carry, ans := 0, ""
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry > 0; {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		if sum >= 2 {
			carry = 1
			sum -= 2
		} else {
			carry = 0
		}
		ans = strconv.Itoa(sum) + ans
	}
	return ans
}
func main() {
	fmt.Println(addBinary("11", "1"))
}
