package main

func numberOfBeams(bank []string) int {
	if len(bank) == 1 {
		return 0
	}

	pre := 0
	result := 0
	for _, r := range bank {
		cur := strings.Count(r, "1")
		if cur > 0 {
			result += pre * cur
			pre = cur
		}
	}
	return result

}
