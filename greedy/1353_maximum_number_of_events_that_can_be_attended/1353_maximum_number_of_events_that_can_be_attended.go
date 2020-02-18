package main

import (
	"fmt"
	"sort"
)

func maxEvents(events [][]int) int {
	var (
		i, j, k, maxEnd int
		length          = len(events)
		visitedDays     []bool
	)
	if length == 0 {
		return 0
	}

	maxEnd = events[0][1]
	sort.Slice(events, func(i, j int) bool {
		if events[i][1] > maxEnd {
			maxEnd = events[i][1]
		}
		if events[i][1] == events[j][1] {
			return events[i][0] > events[j][0]
		}
		return events[i][1] < events[j][1]
	})
	visitedDays = make([]bool, maxEnd+1)

	for i = 0; i < length; i++ {
		for j = events[i][0]; j <= events[i][1]; j++ {
			if !visitedDays[j] {
				visitedDays[j] = true
				break
			}
		}
	}

	for i = 0; i <= maxEnd; i++ {
		if visitedDays[i] {
			k++
		}
	}
	return k
}

func main() {
	events := [][]int{{1, 2}, {2, 3}, {3, 4}}
	res := maxEvents(events)
	fmt.Println(res)
}
