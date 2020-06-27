package main

import (
	"fmt"
	"sort"
)

func average(salary []int) float64 {
	sum := 0
	cnt := 0
	sort.Ints(salary)
	salary = salary[1 : len(salary)-1]
	for _, i := range salary {
		sum += i
		cnt++
	}
	// fmt.Println(sum, cnt)
	return float64(sum) / float64(cnt)
}

func main() {
	salary := []int{4000, 3000, 1000, 2000}
	res := average(salary)
	fmt.Println(res)
}
