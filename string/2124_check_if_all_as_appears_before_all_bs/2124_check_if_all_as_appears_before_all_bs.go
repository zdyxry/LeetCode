package main

func checkString(s string) bool {
	var flag bool
	for _, c := range s {
		if c == 'b' {
			flag = true
		} else {
			if flag {
				return false
			}
		}
	}
	return true
}
