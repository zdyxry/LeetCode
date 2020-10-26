package main

import "fmt"

func slowestKey(releaseTimes []int, keysPressed string) byte {
	t := releaseTimes[0]
	ans := keysPressed[0]
	for i := 1; i < len(releaseTimes); i++ {
		if releaseTimes[i]-releaseTimes[i-1] >= t {
			ans = keysPressed[i]
			t = releaseTimes[i] - releaseTimes[i-1]
		}
	}
	return ans
}

func main() {
	releaseTimes := []int{9, 29, 49, 50}
	keysPressed := "cbcd"
	res := slowestKey(releaseTimes, keysPressed)
	fmt.Println(res)
}
