package main

func areOccurrencesEqual(s string) bool {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}

	current := 0
	for _, v := range m {
		if current == 0 {
			current = v
		} else {
			if current != v {
				return false
			}
		}
	}
	return true
}
