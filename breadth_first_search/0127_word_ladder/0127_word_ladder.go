package main

import (
	"fmt"
)

func ladderLength(beginWord, endWord string, words []string) int {
	dict := make(map[string]bool, len(words))
	for i := 0; i < len(words); i++ {
		dict[words[i]] = true
	}

	delete(dict, beginWord)

	queue := make([]string, 0, len(words))

	var trans func(string) bool
	trans = func(word string) bool {
		bytes := []byte(word)
		for i := 0; i < len(bytes); i++ {
			diffLetter := bytes[i]
			for j := 0; j < 26; j++ {
				b := 'a' + byte(j)
				if b == diffLetter {
					continue
				}

				bytes[i] = b

				if dict[string(bytes)] {
					if string(bytes) == endWord {
						return true
					}

					queue = append(queue, string(bytes))
					delete(dict, string(bytes))
				}
			}

			bytes[i] = diffLetter
		}

		return false
	}

	queue = append(queue, beginWord)
	dist := 1

	for len(queue) > 0 {
		qLen := len(queue)

		for i := 0; i < qLen; i++ {
			word := queue[0]
			queue = queue[1:]

			if trans(word) {
				return dist + 1
			}
		}

		dist++
	}

	return 0
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	words := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	res := ladderLength(beginWord, endWord, words)
	fmt.Println(res)
}
