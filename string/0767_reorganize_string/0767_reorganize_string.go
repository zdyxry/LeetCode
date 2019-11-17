package main 

import (
	"fmt"
	"sort"
)

func reorganizeString(s string) string {
	n := len(s)
	bs := []byte(s)

	count := make([][2]int, 26)
	maxCount := 0
	for _,b := range bs {
		count[b-'a'][0]++
		count[b-'a'][1] = int(b-'a')
		maxCount = max(maxCount, count[b-'a'][0])
	}

	if 2 * maxCount > n+1{
		return ""
	}

	sort.Slice(count, func(i,j int) bool {
		if count[i][0] == count[j][0] {
			return count[i][1] < count[j][1]
		}
		return count[i][0] > count[j][0]
	})

	idx := 0
	for i := 0; i < 26 && count[i][0] > 0; i++ {
		b := byte('a' + count[i][1])

		for count[i][0] > 0 {
			bs[idx] = b
			count[i][0]--
			idx += 2
			if idx >= n {
				idx = 1
			}
		}
	}

	return string(bs)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "aab"
	res := reorganizeString(s)
	fmt.Println(res)
}