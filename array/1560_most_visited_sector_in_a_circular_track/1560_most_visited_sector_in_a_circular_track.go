package main

import "fmt"

func mostVisited(n int, rounds []int) []int {
	const maxn int = 101
	cnt := [maxn]int{}

	cnt[rounds[0]]++
	for i := 1; i < len(rounds); i++ {
		cur := rounds[i-1]
		for cur != rounds[i] {
			cur = (cur + 1) % n
			if cur == 0 {
				cur = n
			}
			cnt[cur]++
		}
	}

	maxCnt := 0
	for i := 1; i <= n; i++ {
		if cnt[i] > maxCnt {
			maxCnt = cnt[i]
		}
	}

	ans := []int{}
	for i := 1; i <= n; i++ {
		if cnt[i] == maxCnt {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	n := 4
	rounds := []int{1, 3, 1, 2}
	res := mostVisited(n, rounds)
	fmt.Println(res)
}
