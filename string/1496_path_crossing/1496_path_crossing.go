package main

import "fmt"

func isPathCrossing(path string) bool {
	visited := make(map[string]bool)

	visited["00"] = true

	x := 0
	y := 0

	for _, val := range path {
		switch val {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		default:
			x--
		}

		key := fmt.Sprintf("%d%d", x, y)

		_, ok := visited[key]

		if ok {
			return true
		}

		visited[key] = true
	}

	return false
}

func main() {
	path := "NES"
	res := isPathCrossing(path)
	fmt.Println(res)
}
