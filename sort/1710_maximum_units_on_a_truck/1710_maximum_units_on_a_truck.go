package main

import (
	"fmt"
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	var res int
	for _, v := range boxTypes {
		if truckSize <= v[0] {
			res += v[1] * truckSize
			break
		} else {
			truckSize -= v[0]
			res += v[0] * v[1]
		}
	}
	return res
}

func main() {
	boxTypes := [][]int{{1, 3}, {2, 2}, {3, 1}}
	res := maximumUnits(boxTypes, 4)
	fmt.Println(res)
}
