package main

import "fmt"

func lemonadeChange(bills []int) bool {
	fifths, tens := 0, 0

	for i := range bills {
		if bills[i] == 5 {
			fifths++
		} else if bills[i] == 10 {
			fifths--
			tens++
		} else if tens > 0 {
			tens--
			fifths--
		} else {
			fifths -= 3
		}
		if fifths < 0 {
			return false
		}
	}
	return true
}

func main() {
	bills := []int{5, 5, 5, 10, 20}
	res := lemonadeChange(bills)
	fmt.Println(res)
}
