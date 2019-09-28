package main 

import "fmt"

func setZeros(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))

	for i:=0; i< len(matrix);i++ {
		for j := 0; j< len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i := 0; i < len(matrix);i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0 
			}
		}
	}
}

func main() {
	matrix := [][]int{{1,1,1}, {1,0,1}, {1,1,1}}

	setZeros(matrix)
	fmt.Println(matrix)
}