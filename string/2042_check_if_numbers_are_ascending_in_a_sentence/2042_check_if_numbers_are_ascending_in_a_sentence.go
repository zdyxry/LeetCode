package main

func areNumbersAscending(s string) bool {
	prev := 0
	for _, token := range strings.Split(s, " ") {
		if unicode.IsDigit(rune(token[0])) {
			v, _ := strconv.Atoi(token)
			if v <= prev {
				return false
			}
			prev = v
		}
	}
	return true
}
