package main

import "fmt"

func getFolderNames(names []string) []string {
	m := make(map[string]int)
	var res []string

	for _, n := range names {
		if m[n] > 0 {
			var max int
			var tmpN string

			for {
				max = m[n]
				m[n]++

				tmpN = fmt.Sprintf("%s(%d)", n, max)
				if m[tmpN] == 0 {
					break
				}
			}

			m[tmpN]++
			res = append(res, tmpN)

		} else {
			m[n]++
			res = append(res, n)
		}
	}
	return res
}

func main() {
	names := []string{"pes", "fifa", "gta", "pes(2019)"}
	res := getFolderNames(names)
	fmt.Println(res)
}
