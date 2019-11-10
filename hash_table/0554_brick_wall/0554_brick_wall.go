package main 

import "fmt"


func leastBricks(wall [][]int) int {
	m := len(wall)

	count := make(map[int]int, m)

	for i := 0; i < m;i++ {
		sum := wall[i][0]
		for j:=1;j<len(wall[i]);j++ {
			count[sum]++
			sum += wall[i][j]
		}
	}

	max := 0
	for _, edges := range count {
		if max < edges {
			max = edges
		}
	}

	return m - max
}


func main() {
	wall := [][]int{{1,2,2,1},{3,1,2},{1,3,2},{2,4},{3,1,2},{1,3,1,1}}
	res := leastBricks(wall)
	fmt.Println(res)
}