package main

import "fmt"

func alphabetBoardPath(target string) string {
	x, y := 0, 0
	moves := make([]byte, 0)

	for _, s := range target {
		n := int(s - 'a')
		nx, ny := n%5, n/5
		dx, dy := nx-x, ny-y

		if y == 5 {
			if dy != 0 {
				moves = append(moves, 'U')
				dy += 1
			}
		}

		for i := 0; i < abs(dx); i++ {
			if dx > 0 {
				moves = append(moves, 'R')
			} else {
				moves = append(moves, 'L')
			}
		}

		for i := 0; i < abs(dy); i++ {
			if dy > 0 {
				moves = append(moves, 'D')
			} else {
				moves = append(moves, 'U')
			}
		}

		moves = append(moves, '!')
		x, y = nx, ny
	}

	return string(moves)
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func main() {
	target := "leet"
	res := alphabetBoardPath(target)
	fmt.Println(res)
}
