package main

import "fmt"

func angleClock(hour int, minutes int) float64 {
    var (
        locaH, locaM, diff float64
    )

    locaM = float64(6 * minutes)
    locaH = float64(30*hour%360) + float64(minutes)*30.0/60.0

    diff = max(locaH, locaM) - min(locaH, locaM)
    if diff > 180.0 {
        diff = 360 - diff
    }
    return diff
}

func max(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}
func min(a, b float64) float64 {
    if a < b {
        return a
    }
    return b
}

func main() {
	hour := 12
	minutes := 30
	res := angleClock(hour, minutes)
	fmt.Println(res)
}
