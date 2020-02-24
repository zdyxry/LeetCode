package main

import (
	"fmt"
)

func numberOfSubstrings(s string) int {
	if len(s) <= 2 {
		return 0
	}

	i, j := 0, 3

	now := []int{0, 0, 0}
	for z := 0; z < 3; z++ {
		now[s[z]-'a']++
	}

	out := 0
	for i < len(s)-2 {
		if !zeroArray(now) {
			out += (len(s) - j + 1)
			now[s[i]-'a']--
			i++
		} else {
			j++
			if j > len(s) {
				break
			}
			now[s[j-1]-'a']++
		}
	}
	return out
}

func numberOfSubstrings2(s string) int {
	ans := 0
	lo := -1
	cnt := []int{0, 0, 0}

	for hi, c := range s {
		cnt[c-'a']++
		for {
			if zeroArray(cnt) {
				break
			} else {
				ans += len(s) - hi
				lo++
				cnt[s[lo]-'a']--
			}
		}
	}
	return ans
}

func zeroArray(nums []int) bool {
	for _, c := range nums {
		if c == 0 {
			return true
		}
	}
	return false
}

func main() {
	s := "abcc"
	res := numberOfSubstrings2(s)
	fmt.Println(res)
}
