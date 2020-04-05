package main 

func countLargestGroup(n int) int {
	var (
		i int
		total []int
		countMap  = map[int]int{}
	)

	for i = 1; i <= n; i++ {
		countMap[count(i)]++
	}
	for _, v := range countMap {
		total = append(total, v)
	}
	sort.Slice(total, func(i, j int) bool {
		return total[i] > total[j]
	})

	for i = 1; i < len(total); i++ {
		if total[i] != total[0] {
			return i
		}
	}
	return len(total)
}

func count(n int) (rst int) {
	for n > 0 {
		rst += (n % 10)
		n /= 10
	}
	return 
}