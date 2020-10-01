package main

import "fmt"

func canConvertString(s string, t string, k int) bool {
	if len(s) != len(t) {
		return false
	}

	seen := make(map[int]int)

	for i := 0; i < len(s); i++ {
		seen[int(t[i]-s[i]+26)%26]++
	}

	cnt := k / 26
	for i := 25; i >= 1; i-- {
		if i == k%26 {
			cnt++
		}

		if seen[i] > cnt {
			return false
		}
	}

	return true
}

func main() {
	s := "input"
	t := "ouput"
	k := 9
	res := canConvertString(s, t, k)
	fmt.Println(res)
}
