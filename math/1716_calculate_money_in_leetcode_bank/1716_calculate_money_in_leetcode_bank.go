package main

import "fmt"

func totalMoney(n int) int {

	week, day := n/7, n%7
	firstWeek := 7 * 8 / 2
	fmt.Println(week, day)
	// n*a1 + d*n*(n-1)/2
	return (week*firstWeek + 7*week*(week-1)/2) + (week+1)*day + day*(day-1)/2
}

func main() {
	n := 10
	res := totalMoney(n)
	fmt.Printf("%d", res)
}
