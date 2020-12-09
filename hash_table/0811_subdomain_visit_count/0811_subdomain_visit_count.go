package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {

	m := make(map[string]int, len(cpdomains)*4)

	for i := range cpdomains {
		cur := strings.Fields(cpdomains[i])
		times, domain := cur[0], cur[1]
		count, _ := strconv.Atoi(times)
		m[domain] += count
		for j := range domain {
			if domain[j] == '.' {
				m[domain[j+1:]] += count
			}
		}
	}

	res := make([]string, 0, len(m))
	for k, v := range m {
		res = append(res, strconv.Itoa(v)+" "+k)
	}

	return res
}

func main() {
	cpdomains := []string{"9001 discuss.leetcode.com"}
	res := subdomainVisits(cpdomains)
	fmt.Println(res)
}
