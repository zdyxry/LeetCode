package main

func checkAlmostEquivalent(word1, word2 string) bool {
	diff := [26]int{} // 统计每个字符的频率差值
	for _, ch := range word1 {
		diff[ch-'a']++
	}
	for _, ch := range word2 {
		diff[ch-'a']--
	}
	for _, d := range diff {
		if d < -3 || d > 3 {
			return false
		}
	}
	return true
}
