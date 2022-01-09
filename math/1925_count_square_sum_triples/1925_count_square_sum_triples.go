package main

func countTriples(n int) int {
	num := 0
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			for c := 1; c <= n; c++ {
				if c*c == a*a+b*b {
					num += 1
				}
			}
		}
	}
	return num
}
