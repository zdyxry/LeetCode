package main

func timeRequiredToBuy(tickets []int, k int) (ans int) {
	for i, t := range tickets {
		if i <= k {
			ans += min(t, tickets[k])
		} else {
			ans += min(t, tickets[k]-1)
		}
	}
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
