package main

func oddCells(n int, m int, indices [][]int) int {
	row := make([]int, n)
	col := make([]int, m)

	for _, indice := range indices {
		row[indice[0]]++
		col[indice[1]]++
	}

	res := 0
	for _, v := range row {
		for _, vv := range col {
			if (v+vv)%2 == 1 {
				res++
			}
		}
	}

	return res
}
