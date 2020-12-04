package main

import "math"

func canEat(piles []int, speed, h int) bool {
	sum := 0.00
	for _, v := range piles {
		sum += math.Ceil(float64(v) * 1.0 / float64(speed))
	}
	return sum > float64(h)
}
func minEatingSpeed(piles []int, H int) int {
	maxVal := 1
	for _, i := range piles {
		maxVal = int(math.Max(float64(maxVal), float64(i)))
	}
	low := 1
	height := maxVal
	for low < height {
		mid := (low + height) / 2
		if canEat(piles, mid, H) {
			low = mid + 1
		} else {
			height = mid
		}
	}
	return low
}

func main() {
	piles := []int{3, 6, 7, 11}
	H := 8
	res := minEatingSpeed(piles, H)
	print(res)
}
