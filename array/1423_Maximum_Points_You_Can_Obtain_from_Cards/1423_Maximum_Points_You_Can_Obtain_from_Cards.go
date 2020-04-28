package main

import "fmt"

func maxScore(cardPoints []int, k int) int {
	var leftSum int
	for i := 0; i < k; i++ {
		leftSum = leftSum + cardPoints[i]
	}
	listLen := len(cardPoints)
	if k == listLen {
		return leftSum
	}
	max := leftSum
	var rightSum int
	for i := 0; i < k; i++ {
		leftSum = leftSum - cardPoints[k-i-1]
		rightSum = rightSum + cardPoints[listLen-i-1]
		if leftSum+rightSum > max {
			max = leftSum + rightSum
		}
	}

	return max
}

func main() {
	cardPoints := []int{1, 2, 3, 4, 5, 6, 1}
	k := 3
	res := maxScore(cardPoints, k)
	fmt.Println(res)
}
