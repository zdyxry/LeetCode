package main

func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")

	count := 0
	for _, w := range words {
		isContain := false
		for _, b := range brokenLetters {
			if strings.ContainsRune(w, b) {
				isContain = true
				break
			}
		}

		if !isContain {
			count++
		}
	}

	return count
}
