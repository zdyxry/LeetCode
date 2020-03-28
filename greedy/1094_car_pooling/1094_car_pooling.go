package main

import "fmt"


func carPooling(trips [][]int, capacity int) bool {
	cost := make([]int, 1001)

	for _, t := range trips {
		cost[t[1]] += t[0]
		cost[t[2]] -= t[0]
	}

	c := 0
	for i := range cost {
		c += cost[i]
		if c > capacity {
			return false
		}
	}
	return true
}

func main() {
	trips := [][]int{{2,1,5},{3,3,7}}
	capacity := 4
	res := carPooling(trips, capacity)
	fmt.Println(res)
}