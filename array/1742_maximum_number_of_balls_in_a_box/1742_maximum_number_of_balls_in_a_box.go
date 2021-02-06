package main

import (
	"fmt"
	"math"
)

func countBalls(lowLimit int, highLimit int) int {
	m := make(map[int]int)
	value := 0
	for i := lowLimit; i <= highLimit; i++ {
		sum := 0
		c := i
		for c != 0 {
			sum += c % 10
			c /= 10
		}
		m[sum]++
		value = int(math.Max(float64(m[sum]), float64(value)))
	}
	return value
}

func main() {
	lowLimit := 5
	highLimit := 15
	res := countBalls(lowLimit, highLimit)
	fmt.Println(res)
}
