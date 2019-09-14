package main 

import "fmt"

func nextPermutation(a []int) {
	left := len(a) - 2
	for 0 <= left && a[left] >= a[left + 1] {
		left--
	}

	reverse(a, left+1)

	if left == -1 {
		return
	}

	right := search(a, left+1, a[left])
	a[left], a[right] = a[right],a[left]
}

func reverse(a []int, l int) {
	r := len(a) - 1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}

func search(a []int, l, target int) int {
	r := len(a) - 1
	l--
	for l+1 < r {
		mid := (l+r) / 2
		if target < a[mid] {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func main() {
	a := []int{1,2,3}
	nextPermutation(a)
	fmt.Println(a)
}