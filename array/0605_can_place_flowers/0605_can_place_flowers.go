package main

import (
	"fmt"
)

func canPlaceFlowers(flowerbed []int, n int) bool {
	l := len(flowerbed)
	for i := 0; i < l; i++ {
		if flowerbed[i] == 0 && // 检查 i 的值
			((i+1 < l && flowerbed[i+1] == 0) || i+1 >= l) && // 检查 i+1 的值
			((i-1 >= 0 && flowerbed[i-1] == 0) || i-1 < 0) { // 检查 i-1 的值
			flowerbed[i] = 1
			n--
		}
	}

	return n <= 0
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 0, 1}
	fmt.Println(canPlaceFlowers(flowerbed, 2))
}
