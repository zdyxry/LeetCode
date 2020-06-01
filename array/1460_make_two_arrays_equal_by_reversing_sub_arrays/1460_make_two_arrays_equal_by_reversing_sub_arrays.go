package main

import "fmt"

func canBeEqual(target []int, arr []int) bool {
	tmp := make(map[int]int)
	for _, v := range target {
		tmp[v]++
	}
	for _, v := range arr {
		if _, ok := tmp[v]; ok {
			tmp[v]--
		} else {
			return false
		}
		if tmp[v] == 0 {
			delete(tmp, v)
		}
	}
	return len(tmp) == 0
}

func main() {
	target := []int{1, 2, 3, 4}
	arr := []int{2, 4, 1, 3}
	res := canBeEqual(target, arr)
	fmt.Println(res)
}
