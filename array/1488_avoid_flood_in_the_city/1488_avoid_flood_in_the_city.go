package main

import "fmt"

func avoidFlood(rains []int) []int {
	var zeros []int
	res := make([]int, len(rains))
	m := make(map[int]int)

	for i, r := range rains {
		if r == 0 {
			zeros = append(zeros, i)
			continue
		}

		if _, ok := m[r]; ok {
			zi := -1

			for ti, tzi := range zeros {
				if tzi > m[r] {
					zi = tzi
					zeros = append(zeros[:ti], zeros[ti+1:]...)
					break
				}
			}

			if zi == -1 {
				return []int{}
			}

			res[zi] = r
		}

		m[r] = i
		res[i] = -1
	}

	for i := range res {
		if res[i] == 0 {
			res[i] = 1
		}
	}
	return res
}

func main() {
	rains := []int{1, 2, 3, 4}
	res := avoidFlood(rains)
	fmt.Println(res)
}
