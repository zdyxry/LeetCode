package main

func getLucky(s string, k int) int {
	b := []byte(s)
	t := ""
	for i := range b {
		b[i] = b[i] - 'a' + 1
		t += strconv.Itoa(int(b[i]))
	}
	b = []byte(t)
	var r int
	for i := 0; i < k; i++ {
		r = 0
		for j := range b {
			r = r + int(b[j]) - '0'
		}
		b = []byte(strconv.Itoa(r))
	}
	return r
}
