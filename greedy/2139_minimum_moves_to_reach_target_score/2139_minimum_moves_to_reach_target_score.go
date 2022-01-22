package main

func minMoves(target int, maxDoubles int) int {
	result := 0

	for target > 1 {
		if maxDoubles > 0 {
			result += 1
			if target%2 == 0 {
				maxDoubles -= 1
				target = target / 2
			} else {
				target -= 1
			}
		} else {
			result += (target - 1)
			target = 1
		}
	}
	return result
}
