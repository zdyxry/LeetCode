package main

func countPrefixes(words []string, s string) int {
	result := 0
	for _, w := range words {
		if strings.HasPrefix(s, w) {
			result += 1
		}
	}
	return result

}
