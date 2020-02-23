1358. Number of Substrings Containing All Three Characters


Medium


Given a string s consisting only of characters a, b and c.

Return the number of substrings containing at least one occurrence of all these characters a, b and c.

 

Example 1:

```
Input: s = "abcabc"
Output: 10
Explanation: The substrings containing at least one occurrence of the characters a, b and c are "abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" and "abc" (again). 
```

Example 2:

```
Input: s = "aaacb"
Output: 3
Explanation: The substrings containing at least one occurrence of the characters a, b and c are "aaacb", "aacb" and "acb". 
```

Example 3:

```
Input: s = "abc"
Output: 1
```
 

Constraints:

3 <= s.length <= 5 x 10^4  
s only consists of a, b or c characters.

## 方法


```go
func numberOfSubstrings(s string) int {
    if len(s) <= 2{
        return 0
    }
    i, j := 0, 3                    // 初始化指针

    now := []int{0,0,0}             // 三长度的数组用于存放当前子串a b c各自数量

    for z:=0; z<3; z++{             // 初始化数量数组
        now[s[z]-'a'] ++
    }

    out := 0
    for i<len(s)-2 {                // 边界条件可以是这个也可以是j <= len(s) 
        if !zeroInArray(now){       // abc都不为0个
            out += (len(s)-j+1)     // 加结果
            now[s[i]-'a'] --        // 取出第i个字符
            i++                     
        } else {                        
            j++
            if j > len(s){break}    // 越界弹出
            now[s[j-1]-'a'] ++      // 增加第j个字符
        }
    }
    return out
}

func zeroInArray(nums []int) bool { // 判断数量数组中是否有0
    for _, c := range nums{
        if c == 0{
            return true
        }
    }
    return false
}
```





```python
class Solution(object):
    def numberOfSubstrings(self, s):
        """
        :type s: str
        :rtype: int
        """
        res = i = 0
        count = {c: 0 for c in 'abc'}
        for j in xrange(len(s)):
            count[s[j]] += 1
            while all(count.values()):
                count[s[i]] -= 1
                i += 1
            res += i
        return res
```


```python
class Solution(object):
    def numberOfSubstrings(self, s):
        """
        :type s: str
        :rtype: int
        """
        ans, lo, cnt = 0, -1, {c : 0 for c in 'abc'}
        for hi, c in enumerate(s):
            cnt[c] += 1
            while all(cnt.values()):
                ans += len(s) - hi
                lo += 1
                cnt[s[lo]] -= 1
        return ans
```