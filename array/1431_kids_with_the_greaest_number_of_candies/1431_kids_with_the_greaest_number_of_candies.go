package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, c := range candies {
		if c > max {
			max = c
		}
	}

	ans := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if max-candies[i] <= extraCandies {
			ans[i] = true
		}
	}
	return ans
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	res := kidsWithCandies(candies, extraCandies)
	fmt.Println(res)
}
