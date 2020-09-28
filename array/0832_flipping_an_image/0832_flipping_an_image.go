package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
	for i, row := range A {
		for j := range row {
			if j > (len(row)-1)/2 {
				break
			}

			tmp := A[i][j]
			A[i][j] = A[i][len(row)-1-j]
			A[i][len(row)-1-j] = tmp

			A[i][j] = A[i][j] ^ 1
			if len(row)-1-j != j {
				A[i][len(row)-1-j] = A[i][len(row)-1-j] ^ 1

			}
		}
	}
	return A
}

func main() {
	A := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	res := flipAndInvertImage(A)
	fmt.Println(res)
}
