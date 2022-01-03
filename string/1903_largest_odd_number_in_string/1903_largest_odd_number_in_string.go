package main

import "fmt"

func largestOddNumber(num string) string {
	n := len(num)
	for i := n - 1; i >= 0; i-- {
		if int(num[i])%2 == 1 {
			return num[0 : i+1]
		}
	}
	return ""

}

func main() {
	num := "52"
	res := largestOddNumber(num)
	fmt.Println(res)
}
