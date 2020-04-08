package main

import "fmt"

func checkOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {
	var dx, dy int
	if x1 > x_center {
		dx = x1 - x_center
	} else if x_center > x2 {
		dx = x_center - x2
	} else {
		dx = 0
	}
	if y1 > y_center {
		dy = y1 - y_center
	} else if y_center > y2 {
		dy = y_center - y2
	} else {
		dy = 0
	}
	return dx*dx+dy*dy <= radius*radius
}

func main() {
	radius := 1
	x_center := 0
	y_center := 0
	x1 := 1
	y1 := -1
	x2 := 3
	y2 := 1
	res := checkOverlap(radius, x_center, y_center, x1, y1, x2, y2)
	fmt.Println(res)
}
