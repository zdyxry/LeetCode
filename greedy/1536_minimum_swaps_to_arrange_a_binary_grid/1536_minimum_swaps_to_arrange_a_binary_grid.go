package main

import "fmt"

func minSwaps(grid [][]int) int {
	row := make([]int, len(grid))
	for i := range grid {
		for j := len(grid) - 1; j >= 0; j-- {
			if grid[i][j] == 0 {
				row[i]++
			} else {
				break
			}
		}
	}

	ans := 0
	for i := 0; i < len(row); i++ {
		last := row[i]
		for j := i; j < len(row); j++ {
			if row[i] >= len(grid)-1-j {
				ans += j - i
				row[j] = last
				break
			}
			if j == len(row)-1 {
				return -1
			}
			tmp := row[j]
			row[j] = last
			last = tmp
		}
	}
	return ans
}

func main() {
	grid := [][]int{{0, 0, 1}, {1, 1, 0}, {1, 0, 0}}
	res := minSwaps(grid)
	fmt.Println(res)
}
