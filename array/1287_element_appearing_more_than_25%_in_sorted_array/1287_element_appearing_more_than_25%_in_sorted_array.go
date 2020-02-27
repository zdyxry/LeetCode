package main

import "fmt"

func findSpecialInteger(arr []int) int {
	length := len(arr)
	cur := arr[0]
	cnt := 0
	for _, v := range arr {
		if v == cur {
			cnt++
			if cnt*4 > length {
				return v
			}
		} else {
			cur = v
			cnt = 1
		}
	}
	return -1

}

func main() {
	arr := []int{1, 2, 2, 6, 6, 6, 6, 7, 10}
	res := findSpecialInteger(arr)
	fmt.Println(res)
}
