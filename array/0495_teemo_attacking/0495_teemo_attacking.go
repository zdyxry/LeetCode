package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}
	res := 0
	for i := 1; i < len(timeSeries); i++ {
		gap := timeSeries[i] - timeSeries[i-1]
		if duration > gap {
			res += gap
		} else {
			res += duration
		}
	}
	return res + duration
}

func main() {
	timeSeries := []int{1, 2}
	duration := 2
	res := findPoisonedDuration(timeSeries, duration)
	fmt.Println(res)
}
