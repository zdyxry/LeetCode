package main

import "fmt"

func minDays(bloomDay []int, m int, k int) int {
	l, r, ans := 1, 1, -1

	for _, v := range bloomDay {
		if v > r {
			r = v
		}
	}

	for l <= r {
		num, mid, cnt := 0, (l+r)/2, 0
		for idx := 0; idx < len(bloomDay); idx++ {
			if bloomDay[idx] <= mid {
				cnt++
				if cnt == k {
					num++
					if num == m {
						break
					}
					cnt = 0
				}
			} else {
				cnt = 0
			}
		}
		if num == m {
			ans = mid
			if l == r {
				break
			}
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans
}

func main() {
	bloomDay := []int{1, 10, 3, 10, 2}
	m := 3
	k := 1
	res := minDays(bloomDay, m, k)
	fmt.Println(res)
}
