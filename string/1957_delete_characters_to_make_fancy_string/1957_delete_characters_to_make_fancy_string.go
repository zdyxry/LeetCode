package main

func makeFancyString(s string) string {
	res := []rune{}
	for _, c := range s {
		if len(res) >= 2 && res[len(res)-1] == c && res[len(res)-2] == c {
			continue
		}
		res = append(res, c)
	}
	return string(res)
}
