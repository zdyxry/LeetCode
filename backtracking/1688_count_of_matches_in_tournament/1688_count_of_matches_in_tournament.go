package main

import "fmt"

func numberOfMatches(n int) int {
	var result int
	for n > 1 {
		if n%2 == 0 {
			result += n / 2
			n = n / 2
		} else {
			result += (n - 1) / 2
			n = (n-1)/2 + 1
		}
	}

	return result
}

func main() {
	n := 14
	res := numberOfMatches(n)
	fmt.Println(res)
}
