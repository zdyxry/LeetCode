package main

import "fmt"

func duplicateZeros(arr []int) {
	for index := 0; index < len(arr); index++ {
		if arr[index] == 0 && index+1 < len(arr) {
			arr = append(arr[:index+1], arr[index:len(arr)-1]...)
			index++
		}
	}
}

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
	fmt.Println(arr)
}
