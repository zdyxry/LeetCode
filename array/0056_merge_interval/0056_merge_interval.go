package main 

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	for i:= 0; i < len(intervals); i++ {
		for j := i+1;j<len(intervals);j++ {
			a := intervals[i]
			b := intervals[j]
			if a[1] >= b[0] && a[0] <= b[1] {
				arr := []int{
					a[0],
					a[1],
					b[0],
					b[1],
				}
				sort.Ints(arr)
				intervals[i] = []int{arr[0], arr[3]}

				intervals = append(intervals[:j], intervals[j+1:]...)
				j = i
			}
		}
	}

	return intervals
}

func main() {
	intervals := [][]int{{1,3},{2,6},{8,10},{15,18}}
	res := merge(intervals)
	fmt.Println(res)
}