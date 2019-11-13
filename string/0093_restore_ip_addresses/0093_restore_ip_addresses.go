package main 

import "fmt"


func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 || n > 12 {
		return []string{}
	}

	res := []string{}
	combination := make([]string, 4)

	var dfs func(int, int)
	dfs = func(idx, begin int) {
		if idx == 3 {
			temp := s[begin:]
			if isOk(temp) {
				combination[3] = temp
				res = append(res, IP(combination))
			}
			return
		}

		maxRemain := 3 * (3 - idx)

		for end := begin + 1; end <= n-(3-idx); end++ {
			if end+maxRemain < n {
				continue
			}

			if end - begin > 3 {
				break
			}

			temp := s[begin:end]
			if isOk(temp) {
				combination[idx] = temp
				dfs(idx+1, end)
			}
		}
	}

	dfs(0, 0)

	return res
}

func IP(s []string) string {
	return fmt.Sprintf("%s.%s.%s.%s", s[0], s[1], s[2], s[3])
}

func isOk(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}

	if len(s) < 3 {
		return true
	}

	switch s[0] {
	case '1':
		return true
	case '2':
		if '0' <= s[1] && s[1] <= '4' {
			return true
		}
		if s[1] == '5' && '0' <= s[2] && s[2] <= '5' {
			return true
		}
	}

	return false
}


func main() {
	s := "25525511135"
	res := restoreIpAddresses(s)
	fmt.Println(res)
}