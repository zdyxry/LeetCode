package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	short := shortest(strs)

	for i, r := range short {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}

	return short
}

func shortest(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for _, s := range strs {
		if len(res) > len(s) {
			res = s
		}
	}

	return res
}
func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
