package main

import "fmt"

func processQueries(queries []int, m int) []int {
	var p = make([]int, m)
	var res = make([]int, len(queries))
	for i := 0; i < m; i++ {
		p[i] = i + 1
	}
	for index, value := range queries {
		//found value in p's index
		var i int
		for i = 0; i < len(p); i++ {
			if p[i] == value {
				res[index] = i
				break
			}
		}
		//move p
		tmp := p[i]
		copy(p[1:i+1], p[0:i])
		p[0] = tmp
	}
	return res
}

func main() {
	queries := []int{3, 1, 2, 1}
	m := 5
	res := processQueries(queries, m)
	fmt.Println(res)
}
