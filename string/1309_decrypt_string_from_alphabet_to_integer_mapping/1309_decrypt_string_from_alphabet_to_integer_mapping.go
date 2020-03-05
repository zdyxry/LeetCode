package main

import "fmt"

func freqAlphabets(s string) string {
	rs := []byte{}
	for i := 0; i < len(s); i++ {
		if i < len(s)-2 && s[i+2] == '#' {
			if s[i] == '1' {
				rs = append(rs, s[i+1]-'0'+'j')
			} else {
				rs = append(rs, s[i+1]-'0'+'t')
			}
			i += 2
		} else {
			rs = append(rs, s[i]-'1'+'a')
		}
	}
	return string(rs)
}

func main() {
	s := "10#11#12"
	res := freqAlphabets(s)
	fmt.Println(res)
}
