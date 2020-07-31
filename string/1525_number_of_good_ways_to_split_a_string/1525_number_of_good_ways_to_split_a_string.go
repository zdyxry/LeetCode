package main

import "fmt"

func numSplits(s string) int {
	freqLeft := [26]int{}
	freqRight := [26]int{}
	for i := 0; i < len(s); i++ {
		freqRight[s[i]-'a']++
	}
	var res int
	for i := 0; i < len(s); i++ {
		freqLeft[s[i]-'a']++
		freqRight[s[i]-'a']--
		if freq(freqLeft) == freq(freqRight) {
			res++
		}
	}
	return res
}

func freq(fre [26]int) int {
	var res int
	for i := 0; i < len(fre); i++ {
		if fre[i] > 0 {
			res++
		}
	}
	return res
}

func main() {
	s := "aacaba"
	res := numSplits(s)
	fmt.Println(res)
}
