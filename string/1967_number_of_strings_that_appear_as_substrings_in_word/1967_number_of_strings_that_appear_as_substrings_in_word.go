package main

func numOfStrings(patterns []string, word string) int {
	result := 0
	for _, p := range patterns {
		//strings.Contains("something", "some")
		if strings.Contains(word, p) {
			result += 1
		}
	}
	return result

}
