package main

import "fmt"

func countGoodTriplets(arr []int, a int, b int, c int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					res += 1
				}
			}
		}
	}
	return res

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	arr := []int{3, 0, 1, 1, 9, 7}
	a := 7
	b := 2
	c := 3
	res := countGoodTriplets(arr, a, b, c)
	fmt.Println(res)
}
