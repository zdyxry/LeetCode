93. Restore IP Addresses


Medium


Given a string containing only digits, restore it by returning all possible valid IP address combinations.

Example:

```
Input: "25525511135"
Output: ["255.255.11.135", "255.255.111.35"]
```


## 方法


```go

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
			if isOK(temp) {
				combination[3] = temp
				res = append(res, IP(combination))
			}
			return
		}

		// 剩余 IP 段，最多需要的宽度
		maxRemain := 3 * (3 - idx)

		for end := begin + 1; end <= n-(3-idx); end++ {
			if end+maxRemain < n {
				// 后面的 IP 段 至少有一个超过了 3 个字符
				// 说明此IP段短了
				continue
			}

			if end-begin > 3 {
				// 此 IP 段长度，超过 3 位了
				break
			}

			temp := s[begin:end]
			if isOK(temp) {
				combination[idx] = temp
				dfs(idx+1, end)
			}
		}
	}

	dfs(0, 0)

	return res
}

// IP 返回 s 代表的 IP 地址
func IP(s []string) string {
	return fmt.Sprintf("%s.%s.%s.%s", s[0], s[1], s[2], s[3])
}

// 由程序其他部分保证了 len(s) <= 3
func isOK(s string) bool {
	// "0"  可以
	// "01" 不可以
	if len(s) > 1 && s[0] == '0' {
		return false
	}

	if len(s) < 3 {
		return true
	}

	// len(s) == 3
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
```



```python
class Solution(object):
    def restoreIpAddresses(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        result = []
        self.restoreIpAddressesRecur(result, s, 0, "", 0)
        return result

    def restoreIpAddressesRecur(self, result, s, start, current, dots):
        # pruning to improve performance
        if (4 - dots) * 3 < len(s) - start or (4 - dots) > len(s) - start:
            return

        if start == len(s) and dots == 4:
            result.append(current[:-1])
        else:
            for i in xrange(start, start + 3):
                if len(s) > i and self.isValid(s[start:i + 1]):
                    current += s[start:i + 1] + '.'
                    self.restoreIpAddressesRecur(result, s, i + 1, current, dots + 1)
                    current = current[:-(i - start + 2)]

    def isValid(self, s):
        if len(s) == 0 or (s[0] == '0' and s != "0"):
            return False
        return int(s) < 256
```