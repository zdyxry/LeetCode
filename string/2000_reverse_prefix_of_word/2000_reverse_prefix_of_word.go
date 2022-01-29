package main

func reversePrefix(word string, ch byte) string {
	index := -1
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			index = i
			break
		}
	}
	if index == -1 {
		return word
	}
	res := make([]byte, 0, len(word))
	for i := index; i >= 0; i-- {
		res = append(res, word[i])
	}
	for i := index + 1; i < len(word); i++ {
		res = append(res, word[i])
	}
	return string(res)
}
