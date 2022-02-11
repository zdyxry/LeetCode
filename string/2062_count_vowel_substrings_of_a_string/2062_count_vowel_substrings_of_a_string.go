package main

func countVowelSubstrings(word string) int {
	var res int
	var cnt [26]int
	n := len(word)
	for i := 0; i < n; i++ {
		cnt = [26]int{}
		for j := i; j < n; j++ {
			if word[j] != 'a' && word[j] != 'e' && word[j] != 'i' && word[j] != 'o' && word[j] != 'u' {
				break
			}
			cnt[word[j]-'a']++
			if cnt['a'-'a'] > 0 && cnt['e'-'a'] > 0 && cnt['i'-'a'] > 0 && cnt['o'-'a'] > 0 && cnt['u'-'a'] > 0 {
				res++
			}
		}
	}
	return res
}
