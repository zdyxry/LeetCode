package main 

import (
	"fmt"
)

func isUgly(num int) bool {
	if num <= 0 {
		return false
	}

	factors := []int{2,3,5}
	for _, factor := range factors{
		for num % factor == 0 {
			num /= factor
		}
	}

	return num == 1
}

func main() {
	num := 6
	res := isUgly(num)
	fmt.Println(res)
}