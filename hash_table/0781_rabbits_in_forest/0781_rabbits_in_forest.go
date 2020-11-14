package main

import "fmt"

func numRabbits(answers []int) int {
	hash := make(map[int]int)
	n := 0
	for _, v := range answers {
		hash[v]++
	}
	// fmt.Println(hash)
	for k, v := range hash {
		k += 1
		n += (v / k) * k
		if (v % k) > 0 {
			n += k
		}
	}
	return n
}

func main() {
	answers := []int{1, 0, 1, 0, 0}
	res := numRabbits(answers)
	fmt.Println(res)
}
