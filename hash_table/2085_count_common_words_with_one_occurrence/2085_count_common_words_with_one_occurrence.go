package main

func countWords(words1 []string, words2 []string) int {
	result := 0
	cnt1 := make(map[string]int)
	cnt2 := make(map[string]int)
	for _, w := range words1 {
		cnt1[w] += 1
	}
	for _, w := range words2 {
		cnt2[w] += 1
	}
	for k, v := range cnt1 {
		if v == 1 && cnt2[k] == 1 {
			result += 1
		}
	}
	return result

}
