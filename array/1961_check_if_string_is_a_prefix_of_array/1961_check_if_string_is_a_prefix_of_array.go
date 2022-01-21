package main

func isPrefixString(s string, words []string) bool {
	comp := ""
	for _, word := range words {
		comp += word
		if comp == s {
			return true
		}
	}
	return false
}
