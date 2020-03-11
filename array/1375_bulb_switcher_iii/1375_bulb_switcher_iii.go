package main

import "fmt"

func numTimesAllBlue(light []int) int {
	var (
		i       int
		rst     int
		maxOpen = 0
		length  = len(light)
	)

	for i = 0; i < length; i++ {
		if light[i] > maxOpen {
			maxOpen = light[i]
		}
		if maxOpen == i+1 {
			rst++
		}
	}
	return rst
}

func main() {
	light := []int{2, 1, 3, 5, 4}
	res := numTimesAllBlue(light)
	fmt.Println(res)
}
