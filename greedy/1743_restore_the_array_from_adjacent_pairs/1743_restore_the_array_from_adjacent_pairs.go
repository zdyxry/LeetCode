package main

import "fmt"

func restoreArray(adjacentPairs [][]int) []int {

	map1 := map[int][]int{}

	for _, i := range adjacentPairs {
		map1[i[0]] = append(map1[i[0]], i[1])
		map1[i[1]] = append(map1[i[1]], i[0])
	}

	head := 0

	for k, v := range map1 {
		if len(v) == 1 {
			head = k
			break
		}
	}

	//fmt.Println("head", head)
	res := make([]int, 0)

	res = append(res, head)
	set := map[int]int{}
	set[head] = 0
	for len(res) < len(adjacentPairs)+1 {
		for _, i := range map1[head] {
			_, ok := set[i]
			if !ok {
				head = i
				res = append(res, i)
				set[head] = 1
				break
			}
		}
	}

	return res
}

func main() {
	adjacentPairs := [][]int{{2, 1}, {3, 4}, {2, 3}}
	res := restoreArray(adjacentPairs)
	fmt.Println(res)
}
