package main

func makeFancyString(s string) string {
	res := []string{}
	for _, c := range s {
		if len(res) >= 2 && res[len(res)-1] == string(c) && res[len(res)-2] == string(c) {
			continue
		}
		res = append(res, string(c))
	}
	return strings.Join(res, "")
}
