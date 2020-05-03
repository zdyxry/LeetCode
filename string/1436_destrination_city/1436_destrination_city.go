package main

import "fmt"

func destCity(paths [][]string) string {
	m := make(map[string]bool)
	for _, p := range paths {
		m[p[0]] = true
	}
	for _, p := range paths {
		if !m[p[1]] {
			return p[1]
		}
	}
	return ""
}

func main() {
	paths := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	res := destCity(paths)
	fmt.Println(res)
}
