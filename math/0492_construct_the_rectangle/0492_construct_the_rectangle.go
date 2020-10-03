package main

import (
	"fmt"
	"math"
)

func constructRectangle(area int) []int {
	mid := int(math.Sqrt(float64(area)))
	for area%mid != 0 {
		mid--
	}
	return []int{area / mid, mid}
}

func main() {
	area := 1
	res := constructRectangle(area)
	fmt.Println(res)
}
