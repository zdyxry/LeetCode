package main

import "fmt"

func maximumTime(time string) string {
	times := []byte(time)
	if times[0] == '?' {
		if times[1] == '?' || times[1] < '4' {
			times[0] = '2'
		} else {
			times[0] = '1'
		}
	}
	if times[1] == '?' {
		if times[0] == '2' {
			times[1] = '3'
		} else {
			times[1] = '9'
		}
	}
	if times[3] == '?' {
		times[3] = '5'
	}
	if times[4] == '?' {
		times[4] = '9'
	}
	return string(times)
}

func main() {
	time := "2?:?0"
	res := maximumTime(time)
	fmt.Println(res)
}
