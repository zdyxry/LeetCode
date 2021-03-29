1796. Second Largest Digit in a String


Easy


Given an alphanumeric string s, return the second largest numerical digit that appears in s, or -1 if it does not exist.

An alphanumeric string is a string consisting of lowercase English letters and digits.

 

Example 1:

```
Input: s = "dfa12321afd"
Output: 2
Explanation: The digits that appear in s are [1, 2, 3]. The second largest digit is 2.
```

Example 2:

```
Input: s = "abc1111"
Output: -1
Explanation: The digits that appear in s are [1]. There is no second largest digit. 
```

Constraints:

1 <= s.length <= 500   
s consists of only lowercase English letters and/or digits.   


## 方法



```go
func secondHighest(s string) int {
    var counts [10]int
    for _,ch:=range s{
        if ch>='0'&&ch<='9'{
            counts[ch-'0']=1
        }
    }
    sign := false
    for i:=9;i>=0;i--{
        if counts[i]==1 {
            if sign {
                return i
            }
            sign=true
        }
    }
    return -1
}
```


```python
class Solution:
    def secondHighest(self, s: str) -> int:
        q = []
        for ch in s:
            if ch.isdigit():
                n = int(ch)
                if not q or -n != q[0]:
                    heapq.heappush(q, -n)
        if not q: return -1
        heapq.heappop(q)
        return -heapq.heappop(q) if q else -1
```