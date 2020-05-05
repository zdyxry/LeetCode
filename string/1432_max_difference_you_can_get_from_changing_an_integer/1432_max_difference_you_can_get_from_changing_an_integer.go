package main

import "fmt"

func toDigits(num int) []int {
	if num < 10 {
		return []int{num}
	}
	return append(toDigits(num/10), num%10)
}

func maxDiff(num int) int {
	digits := toDigits(num)

	digitA, digitB := -1, -1
	var zero bool
	for _, d := range digits {
		if digitA < 0 && d < 9 {
			digitA = d
		}
		if digitB < 0 && d >= 1 {
			if d == digits[0] {
				if d > 1 {
					digitB = d
				}
			} else {
				digitB = d
				zero = true
			}
		}
		if digitA >= 0 && digitB >= 0 {
			break
		}
	}
	a, b := 0, 0
	for _, d := range digits {
		a *= 10
		b *= 10
		if d == digitA {
			a += 9
		} else {
			a += d
		}
		if d == digitB {
			if !zero {
				b++
			}
		} else {
			b += d
		}
	}
	return a - b

}

func main() {
	num := 555
	res := maxDiff(num)
	fmt.Println(res)
}
