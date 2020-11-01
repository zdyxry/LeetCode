package main

import "fmt"

func canFormArray(arr []int, pieces [][]int) bool {
	m := make(map[int][]int)

	for _, p := range pieces {
		m[p[0]] = p[1:]
	}

	for i := 0; i < len(arr); {
		piece, ok := m[arr[i]]

		if !ok {
			return false
		}

		i++
		for _, p := range piece {
			if p != arr[i] {
				return false
			}
			i++
		}
	}

	return true
}

func main() {
	arr := []int{85}
	pieces := [][]int{{85}}
	res := canFormArray(arr, pieces)
	fmt.Println(res)
}
