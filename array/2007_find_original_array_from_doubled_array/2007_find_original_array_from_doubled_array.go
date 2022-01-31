package main

func findOriginalArray(changed []int) (original []int) {
	sort.Ints(changed)
	cnt := map[int]int{}
	for _, v := range changed {
		if cnt[v] == 0 {
			cnt[v*2]++
			original = append(original, v)
		} else {
			cnt[v]--
			if cnt[v] == 0 {
				delete(cnt, v)
			}
		}
	}
	if len(cnt) == 0 {
		return
	}
	return nil
}
