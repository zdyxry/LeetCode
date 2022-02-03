package main

func sumOfBeauties(nums []int) int {

	total := 0

	n := len(nums)

	pre := make([]int, n)
	suf := make([]int, n)

	pre[0] = nums[0]
	suf[n-1] = nums[n-1]

	for i := 1; i < n; i++ {
		if pre[i-1] < nums[i] {
			pre[i] = nums[i]
		} else {
			pre[i] = pre[i-1]
		}
	}

	for i := n - 2; i >= 0; i-- {
		if nums[i] < suf[i+1] {
			suf[i] = nums[i]
		} else {
			suf[i] = suf[i+1]
		}
	}

	fmt.Println(pre)
	fmt.Println(suf)

	for i := 1; i <= n-2; i++ {
		if pre[i-1] < nums[i] && nums[i] < suf[i+1] {
			total += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			total++
		}
	}

	return total
}
