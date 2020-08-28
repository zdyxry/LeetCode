package main

import "fmt"

func tictactoe(moves [][]int) string {
	slc := []int{7, 56, 73, 84, 146, 273, 292, 448}
	na, nb := 0, 0
	for i := range moves {
		cur := 1 << (moves[i][0]*3 + moves[i][1])
		if i&1 == 0 {
			na ^= cur
			for _, v := range slc {
				if na&v == v {
					return "A"
				}
			}
		} else {
			nb ^= cur
			for _, v := range slc {
				if nb&v == v {
					return "B"
				}
			}
		}
	}
	if len(moves) == 0 {
		return "Draw"
	}
	return "Pending"
}

func main() {
	moves := [][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}
	res := tictactoe(moves)
	fmt.Println(res)
}
