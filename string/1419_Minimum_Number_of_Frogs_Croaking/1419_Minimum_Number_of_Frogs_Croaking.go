package main

import "fmt"

func minNumberOfFrogs(croakOfFrogs string) int {
	c, r, o, a, k, frog := 0, 0, 0, 0, 0, 0
	for _, v := range croakOfFrogs {
		switch v {
		case 'c':
			c++
			if c > frog {
				frog = c
			}
		case 'r':
			r++
			if r > c {
				return -1
			}
		case 'o':
			o++
			if o > r {
				return -1
			}
		case 'a':
			a++
			if a > o {
				return -1
			}
		case 'k':
			k++
			if k > a {
				return -1
			}
			c--
			r--
			o--
			a--
			k--
		}
	}
	if c+r+o+a+k != 0 {
		return -1
	}
	return frog
}

func main() {
	croakOfFrogs := "croakcroak"
	res := minNumberOfFrogs(croakOfFrogs)
	fmt.Println(res)
}
