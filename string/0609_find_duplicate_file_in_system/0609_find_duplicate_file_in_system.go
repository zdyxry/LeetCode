package main

import (
	"fmt"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	out := [][]string{}
	m := map[string][]string{}
	for _, path := range paths {
		es := strings.Split(path, " ")
		for _, e := range es[1:] {
			i := strings.LastIndexByte(e, '(')
			k := e[i+1 : len(e)-1]
			m[k] = append(m[k], es[0]+"/"+e[:i])
		}
	}
	for _, sl := range m {
		if len(sl) > 1 {
			out = append(out, sl)
		}
	}
	return out
}

func main() {
	paths := []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}
	res := findDuplicate(paths)
	fmt.Println(res)
}
