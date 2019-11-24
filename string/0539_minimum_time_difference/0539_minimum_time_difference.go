package main 

import (
	"fmt"
	"strconv"
	"sort"
)

func findMinDifference(timePoints []string) int {
	n := len(timePoints)
	sort.Strings(timePoints)

	minDiff := 24 * 60 - diff(timePoints[0], timePoints[n-1])
	for i := 1; i < n; i++ {
		minDiff = min(minDiff, diff(timePoints[i-1], timePoints[i]))
	}
	return minDiff
}

func diff(a, b string) int {
	ha, ma := getHourAndMintueOf(a)
	hb, mb := getHourAndMintueOf(b)
	return (hb-ha) * 60 + (mb-ma)
}

func getHourAndMintueOf(s string) (hour, mintue int) {
	hour, _ = strconv.Atoi(s[:2])
	mintue, _ = strconv.Atoi(s[3:])
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	timePoints := []string{"23:59","00:00"}
	res := findMinDifference(timePoints)
	fmt.Println(res)
}