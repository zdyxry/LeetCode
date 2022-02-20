package main

func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}
func isPalindrome(word string) bool {
	n := len(word)
	l, r := 0, n-1
	for l < r {
		if word[l] != word[r] {
			return false
		} else {
			l++
			r--
		}
	}
	return true
}
