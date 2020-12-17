package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(S string, K int) string {
	str := strings.ReplaceAll(S, "-", "")
	for i := len(str) - K; i > 0; i -= K {
		str = str[:i] + "-" + str[i:]
	}
	return strings.ToUpper(str)
}

func main() {
	S := "5F3Z-2e-9-w"
	K := 4
	res := licenseKeyFormatting(S, K)
	fmt.Println(res)
}
