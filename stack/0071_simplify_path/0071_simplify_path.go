package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	lp := len(path)
	stack := make([]string, 0, lp/2)
	dir := make([]byte, 0, lp)

	for i := 0; i < lp; i++ {
		dir = dir[:0]
		for i < lp && path[i] != '/' {
			dir = append(dir, path[i])
			i++
		}

		s := string(dir)

		switch s {
		case ".", "":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, s)
		}

	}
	return "/" + strings.Join(stack, "/")
}

func main() {
	path := "/home/"
	res := simplifyPath(path)
	fmt.Println(res)
}
