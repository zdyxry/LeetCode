package main

func largestInteger(num int) int {
	s := []byte(strconv.Itoa(num))
	var a []byte
	var b []byte
	for i := range s {
		if (s[i]-'0')&1 == 1 {
			a = append(a, s[i])
		} else {
			b = append(b, s[i])
		}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	for i := range s {
		if (s[i]-'0')&1 == 1 {
			s[i] = a[0]
			a = a[1:]
		} else {
			s[i] = b[0]
			b = b[1:]
		}
	}
	r, _ := strconv.Atoi(string(s))
	return r
}
