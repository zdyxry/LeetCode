package main

import "fmt"

func averageWaitingTime(customers [][]int) float64 {
	var wait int
	cur := customers[0][0]

	for _, c := range customers {
		wait += c[1]
		if cur <= c[0] {
			cur = c[0] + c[1]
		} else {
			wait += cur - c[0]
			cur += c[1]
		}
	}

	return float64(wait) / float64(len(customers))
}

func main() {
	customers := [][]int{{1, 2}, {2, 5}, {4, 3}}
	res := averageWaitingTime(customers)
	fmt.Println(res)
}
