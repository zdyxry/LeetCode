package main

import "fmt"

func minOperations(boxes string) []int {
	res := []int{}
	n := len(boxes)
	var total, now int
	for i := 0; i < n; i++ {
		total += now
		res = append(res, total)
		if boxes[i] == '1' {
			now++
		}
	}
	total, now = 0, 0
	for i := n - 1; i >= 0; i-- {
		total += now
		res[i] += total
		if boxes[i] == '1' {
			now++
		}
	}
	return res

}

func main() {
	boxes := "110"
	res := minOperations(boxes)
	fmt.Println(res)
}
