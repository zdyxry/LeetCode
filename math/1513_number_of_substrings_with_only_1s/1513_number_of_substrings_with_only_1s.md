1513. Number of Substrings With Only 1s


Medium


Given a binary string s (a string consisting only of '0' and '1's).

Return the number of substrings with all characters 1's.

Since the answer may be too large, return it modulo 10^9 + 7.

 

Example 1:

```
Input: s = "0110111"
Output: 9
Explanation: There are 9 substring in total with only 1's characters.
"1" -> 5 times.
"11" -> 3 times.
"111" -> 1 time.
```

Example 2:

```
Input: s = "101"
Output: 2
Explanation: Substring "1" is shown 2 times in s.
```

Example 3:

```
Input: s = "111111"
Output: 21
Explanation: Each substring contains only 1's characters.
```

Example 4:

```
Input: s = "000"
Output: 0
```
 

Constraints:
 
s[i] == '0' or s[i] == '1'  
1 <= s.length <= 10^5

## 方法

```go
const mod = 1000000007
func numSub(s string) int {
	sum, count := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count++
		} else {
			sum = (sum + (count+1)*count/2) % mod
			count = 0
		}
	}
	if count != 0 { // 这里注意一下，如果是‘1’结尾的字符串，最后count不为零，需要再算一次
		sum = (sum + (count+1)*count/2) % mod
	}
	return sum
}

```

```python
class Solution:
    def numSub(self, s: str) -> int:
        cnt = 0
        res = 0
        for i in s:
            if i == '1':
                cnt += 1
            else:
                res += cnt * (cnt +1) /2
                cnt = 0
        if cnt > 0:
            res += cnt * (cnt + 1)/ 2
        res %= 1e9+7
        return int(res)
```