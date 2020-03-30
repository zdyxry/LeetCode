package main

import "fmt"

func numTeams(rating []int) int {
	rLen := len(rating)
	count := 0
	for i := 1; i < rLen-1; i++ {
		count1, count2 := 0, 0
		for j := 0; j < i; j++ {
			if rating[j] < rating[i] {
				count1++
			}
		}
		for j := i + 1; j < rLen; j++ {
			if rating[j] < rating[i] {
				count2++
			}
		}
		count += count1*(rLen-1-i-count2) + (i-count1)*count2
	}
	return count
}

func main() {
	rating := []int{2, 5, 3, 4, 1}
	res := numTeams(rating)
	fmt.Println(res)
}
