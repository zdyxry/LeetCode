package main

func numTilePossibilities(tiles string) int {
	used := make([]bool, len(tiles))
	m := make(map[string]struct{})
	backtrack("", used, tiles, m)
	return len(m)
}

func backtrack(now string, used []bool, str string, m map[string]struct{}) {
	if len(now) > 0 {
		m[now] = struct{}{}
	}
	if len(now) == len(str) {
		return
	}
	for i := 0; i < len(str); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		backtrack(now+string(str[i]), used, str, m)
		used[i] = false
	}
}

func main() {
	tiles := "AAB"
	res := numTilePossibilities(tiles)
	print(res)
}
