package main 

func longestBeautifulSubstring(s string) (ans int) {
	const vowel = "aeiou"
	cur, sum := 0, 0
	for i, n := 0, len(s); i < n; {
		start := int
		ch := s[start]
		for i < n && s[i] == ch {
			i++
		}

		if ch != vowel[cur] {
			cur, sum = 0, 0
			if ch != vowel[0] {
				continue
			}
		}

		sum += i - start
		cur++
		if cur == 5 {
			if sum > ans {
				ans = sum
			}
			cur, sum = 0, 0

		}
	}
	return ans
}