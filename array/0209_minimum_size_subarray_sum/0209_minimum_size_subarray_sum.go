package main 

import "fmt"

func minSubArrayLen(s int, a []int) int {
	n := len(a)
	res, i, j, sum := n+1, 0, 0, 0

	for j < n {
		sum += a[j]
		j++

		for sum >= s {
			sum -= a[i]
			i++
			if res > j-i+1 {
				res = j-i+1
			}
		}
	}

	return res % (n+1)
}

func main() {
	a := []int{2,3,1,2,4,3}
	s := 7
	res := minSubArrayLen(s, a)
	fmt.Println(res)
}