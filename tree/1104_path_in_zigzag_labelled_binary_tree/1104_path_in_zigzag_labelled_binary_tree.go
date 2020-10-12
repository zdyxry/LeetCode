package main

import "fmt"

func pathInZigZagTree(label int) []int {
	res := []int{}
	for label > 0 {
		res = append([]int{label}, res...)
		l2 := log2(label)
		label = l2 - 1 - (label-l2)/2
	}
	return res
}

func log2(x int) int {
	for i := 1; i < 31; i++ {
		if x < 1<<i {
			return 1 << (i - 1)
		}
	}
	return -1
}

func main() {
	label := 14
	res := pathInZigZagTree(label)
	fmt.Println(res)
}
