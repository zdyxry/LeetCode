package main 

import "fmt"
import "math"


func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	res := math.MinInt32
	var dfs func(idx int, need int) 
	dfs = func(idx int, need int) {
		if idx == len(toppingCosts) {
			if abs(res - target) > abs(need - target) {
				res = need
			} else if abs(res - target) == abs(need-target) {
				res = min(res, need)
			}
			return
		}
		dfs(idx + 1, need)
		dfs(idx + 1, need + toppingCosts[idx])
		dfs(idx + 1, need + toppingCosts[idx] * 2)
	}

	for i := 0; i < len(baseCosts); i++ {
		dfs(0, baseCosts[i])
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	baseCosts := []int{1,7}
	toppingCosts := []int{3,4}
	target := 10
	res := closestCost(baseCosts, toppingCosts, target)
	fmt.Println(res)
}