package main

import "math"

func nearestValidPoint(x int, y int, points [][]int) int {
	minIndex := -1
	minDist := float64(-1)
	for i := 0; i < len(points); i++ {
		point := points[i]
		if point[0] == x || point[1] == y {
			dist := ManhattanDist(x, y, point[0], point[1])
			if minDist == -1 || minDist > dist {
				minDist = dist
				minIndex = i
			}
		}
	}
	return minIndex
}
func ManhattanDist(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}
