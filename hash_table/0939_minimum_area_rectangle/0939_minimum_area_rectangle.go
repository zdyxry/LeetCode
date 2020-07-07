package main

import "fmt"

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func minAreaRect(points [][]int) int {
	set := map[[2]int]bool{}
	for _, point := range points {
		set[[2]int{point[0], point[1]}] = true
	}
	best := 0

	for i := 0; i < len(points); i++ {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			if p1[0] == p2[0] || p1[1] == p2[1] {
				continue
			}
			p3 := [2]int{p1[0], p2[1]}
			p4 := [2]int{p2[0], p1[1]}
			if !set[p3] || !set[p4] {
				continue
			}
			area := abs(p1[0]-p2[0]) * abs(p1[1]-p2[1])
			if best <= 0 || area < best {
				best = area
			}
		}
	}
	return best
}

func main() {
	points := [][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}}
	res := minAreaRect(points)
	fmt.Println(res)
}
