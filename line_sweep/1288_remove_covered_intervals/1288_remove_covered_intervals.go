package main

import (
	"fmt"
	"sort"
)

func removeCoveredIntervals(intervals [][]int) int {
	if len(intervals) == 0 || len(intervals[0]) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] > intervals[j][1]
		}
	})

	result := 0
	max := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] == intervals[i-1][0] {
			result++
			continue
		}
		if max < intervals[i][1] {
			max = intervals[i][1]
		} else {
			result++
		}
	}
	return len(intervals) - result
}

func main() {
	intervals := [][]int{{1, 4}, {3, 6}, {2, 8}}
	res := removeCoveredIntervals(intervals)
	fmt.Println(res)
}
