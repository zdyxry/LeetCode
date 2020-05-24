package main

import "fmt"

func peopleIndexes(favoriteCompanies [][]string) []int {
	ans := make([]int, 0)
	dist := make(map[int][]map[string]int)
	for _, v := range favoriteCompanies {
		// v - []string
		m := make(map[string]int)
		for _, val := range v {
			// val - company name str
			m[val]++
		}
		dist[len(v)] = append(dist[len(v)], m)
	}
	for k, v := range favoriteCompanies {
		if subset(v, dist) {
			ans = append(ans, k)
		}
	}
	return ans
}

func subset(d1 []string, d2 map[int][]map[string]int) bool {
	ans := 0
	for k, v := range d2 {
		if k >= len(d1) {
			for _, m := range v {
				sub := true
				for _, s := range d1 {
					if _, exist := m[s]; !exist {
						sub = false
						break
					}
				}
				if sub {
					ans++
				}
			}
		}
	}
	if ans >= 2 {
		return false
	} else {
		return true
	}
}

func main() {
	favoriteCompanies := [][]string{{"leetcode", "google", "facebook"}, {"google", "microsoft"}, {"google", "facebook"}, {"google"}, {"amazon"}}
	res := peopleIndexes(favoriteCompanies)
	fmt.Println(res)
}
