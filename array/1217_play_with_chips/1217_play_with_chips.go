package main

import "fmt"

func minCostToMoveChips(chips []int) int {
	if len(chips) == 0 {
		return 0
	}

	oddNums := 0
	evenNums := 0

	for _, k := range chips {
		if k%2 == 0 {
			oddNums += 1
		} else {
			evenNums += 1
		}
	}

	if oddNums > evenNums {
		return evenNums
	} else {
		return oddNums
	}
}

func main() {
	chips := []int{1, 2, 3}
	res := minCostToMoveChips(chips)
	fmt.Println(res)
}
