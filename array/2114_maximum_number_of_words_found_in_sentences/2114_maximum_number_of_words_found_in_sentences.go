package main

func mostWordsFound(sentences []string) int {
	var res int
	for _, sentence := range sentences {
		res = max(res, len(strings.Fields(sentence)))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
