package main

import (
	"fmt"
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	sort.Sort(sort.StringSlice(folder))
	out := []string{}
	out = append(out, folder[0])
	t := folder[0]
	for i := 1; i < len(folder); i++ {
		if strings.HasPrefix(folder[i], t+"/") == false {
			out = append(out, folder[i])
			t = folder[i]
		}
	}
	return out
}

func main() {
	folder := []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}
	res := removeSubfolders(folder)
	fmt.Println(res)
}
