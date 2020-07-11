package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i

}

func main() {
	g := []int{10, 9, 8, 7}
	s := []int{5, 6, 7, 8}
	res := findContentChildren(g, s)
	fmt.Println(res)
}
