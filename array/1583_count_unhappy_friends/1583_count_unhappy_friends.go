package main 


func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	dd := make(map[int][]int)

	for _, pair := range pairs {
		i, x := pair[0], pair[1]
		for idx, j := range preferences[i] {
			if j == x {
				dd[i] = preferences[i][:idx]
			}
		}
		for idx, j := range preferences[x] {
			if j == i {
				dd[x] = preferences[x][:idx]
			}
		}
	}

	ans := 0
    fmt.Println(dd)

	for i := range dd {
        out:
		for _, x := range dd[i] {
			for _, j := range dd[x] {
				if i == j {
					ans++
					break out
				}
			}
		}
	}
	return ans
}