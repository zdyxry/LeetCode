2124. Check if All A's Appears Before All B's


Easy


Given a string s consisting of only the characters 'a' and 'b', return true if every 'a' appears before every 'b' in the string. Otherwise, return false.

 

Example 1:

```
Input: s = "aaabbb"
Output: true
Explanation:
The 'a's are at indices 0, 1, and 2, while the 'b's are at indices 3, 4, and 5.
Hence, every 'a' appears before every 'b' and we return true.
```

Example 2:

```
Input: s = "abab"
Output: false
Explanation:
There is an 'a' at index 2 and a 'b' at index 1.
Hence, not every 'a' appears before every 'b' and we return false.
```

Example 3:
```
Input: s = "bbb"
Output: true
Explanation:
There are no 'a's, hence, every 'a' appears before every 'b' and we return true.
```

Constraints:

1 <= s.length <= 100   
s[i] is either 'a' or 'b'.   


## 方法



```
func checkString(s string) bool {
    var flag bool
	for _, c := range s {
		if c == 'b' {
			flag = true
		} else {
			if flag {
				return false
			}
		}
	}
	return true
}
```


```
class Solution:
    def checkString(self, s: str) -> bool:
        n = len(s)
        lasta = -1   # 'a' 最后一次出现的下标
        firstb = n   # 'b' 第一次出现的下标
        for i in range(n):
            if s[i] == 'a':
                lasta = max(lasta, i)
            else:
                firstb = min(firstb, i)
        return lasta < firstb

```