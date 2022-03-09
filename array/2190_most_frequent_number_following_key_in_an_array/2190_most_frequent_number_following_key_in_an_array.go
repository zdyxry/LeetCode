package main

func mostFrequent(nums []int, key int) (result int) {
	l := len(nums)
	m := make(map[int]int)
	for i, n := range nums {
		if n == key {
			if i+1 < l {
				m[nums[i+1]] += 1
			}
		}
	}
	//fmt.Println(m)
	max := 0
	for k, v := range m {
		if v > max {
			result = k
			max = v
		}
	}

	return result
}
