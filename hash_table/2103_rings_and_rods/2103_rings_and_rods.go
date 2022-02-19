package main

var bits = []int{'B': 1, 'G': 2, 'R': 4}

func countPoints(s string) (ans int) {
	masks := [10]int{}
	for i := 0; i < len(s); i += 2 {
		masks[s[i+1]-'0'] |= bits[s[i]]
	}
	for _, m := range masks {
		if m == 7 { // 7 = 1 + 2 + 4
			ans++
		}
	}
	return
}
