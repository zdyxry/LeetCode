package main

func digitSum(s string, k int) string {
	for len(s) < k {
		var build strings.Builder
		n := len(s)
		for i := 0; i < n; i += k {
			var sum int
			for j := 0; j < k && i+j < n; j++ {
				sum += int(s[i+j] - '0')
			}
			build.WriteString(strconv.Itoa(sum))
		}
		s = build.String()
	}
	return s
}
