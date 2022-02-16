package main

func findEvenNumbers(digits []int) []int {
	f := make([]int, 10)
	for _, n := range digits {
		f[n]++
	}

	r := make([]int, 0)
	for a := 1; a <= 9; a++ {
		if f[a] == 0 {
			continue
		}
		f[a]--

		for b := 0; b <= 9; b++ {
			if f[b] == 0 {
				continue
			}

			f[b]--
			for c := 0; c <= 8; c += 2 {
				if f[c] > 0 {
					r = append(r, 100*a+10*b+c)
				}
			}
			f[b]++
		}
		f[a]++
	}
	return r
}
