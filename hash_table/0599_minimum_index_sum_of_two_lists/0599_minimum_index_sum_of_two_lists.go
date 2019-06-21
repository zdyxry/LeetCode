package main

import (
	"fmt"
)

func findRestaurant(list1 []string, list2 []string) []string {
	if len(list1) > len(list2) {
		list1, list2 = list2, list1
	}

	m2 := make(map[string]int, len(list2))
	for i := range list2 {
		m2[list2[i]] = i
	}

	min := 2000

	res := make([]string, 0, 1000)
	for i, r := range list1 {
		if j, ok := m2[r]; ok {
			if min == i+j {
				res = append(res, r)
			}
			if min > i+j {
				res = append(res[len(res):], r)
				min = i + j
			}
		}
	}
	return res
}

func main() {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	fmt.Println(findRestaurant(list1, list2))
}
