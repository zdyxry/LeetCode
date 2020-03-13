package main

import "fmt"

func numOfMinutes(n int, headID int, managers []int, informTime []int) int {
	subordinates := make([][]int, n)
	for employee, manager := range managers {
		if employee == headID {
			continue
		}
		subordinates[manager] = append(subordinates[manager], employee)
	}

	var dfs func(int) int
	dfs = func(id int) int {
		if len(subordinates[id]) == 0 {
			return 0
		}
		var maxTime int
		for _, subordinate := range subordinates[id] {
			maxTime = max(maxTime, dfs(subordinate))
		}
		return maxTime + informTime[id]
	}
	return dfs(headID)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := 7
	headID := 6
	managers := []int{1, 2, 3, 4, 5, 6, -1}
	informTime := []int{0, 6, 5, 4, 3, 2, 1}
	res := numOfMinutes(n, headID, managers, informTime)
	fmt.Println(res)
}
