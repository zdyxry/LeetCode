package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
	sMap := make(map[string]int, len(s))
	for _, c := range s {
		sMap[string(c)]++
	}
	chars := []string{}
	for c, _ := range sMap {
		chars = append(chars, string(c))
	}
	sort.Strings(chars)
	res := ""
	for len(sMap) > 0 {
		for i := range chars {
			if sMap[chars[i]] > 0 {
				res += string(chars[i])
				sMap[chars[i]]--
			} else {
				delete(sMap, chars[i])
			}
		}
		for i := range chars {
			if sMap[chars[len(chars)-i-1]] > 0 {
				res += string(chars[len(chars)-i-1])
				sMap[chars[len(chars)-i-1]]--
			} else {
				delete(sMap, chars[len(chars)-i-1])
			}
		}
	}
	return res
}

func main() {
	s := "aaaabbbbcccc"
	res := sortString(s)
	fmt.Println(res)
}
