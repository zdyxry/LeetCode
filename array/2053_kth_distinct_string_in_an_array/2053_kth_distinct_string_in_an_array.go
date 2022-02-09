package main

func kthDistinct(arr []string, k int) string {
	tmp := make(map[string]int, len(arr))
	for i := range arr {
		tmp[arr[i]] += 1
	}
	for i := range arr {
		if tmp[arr[i]] == 1 {
			k -= 1
		}
		if k == 0 {
			return arr[i]
		}

	}
	return ""

}

