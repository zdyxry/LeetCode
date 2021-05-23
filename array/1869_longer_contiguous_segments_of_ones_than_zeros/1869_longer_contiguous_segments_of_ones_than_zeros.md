1869. Longer Contiguous Segments of Ones than Zeros


Easy


Given a binary string s, return true if the longest contiguous segment of 1s is strictly longer than the longest contiguous segment of 0s in s. Return false otherwise.

For example, in s = "110100010" the longest contiguous segment of 1s has length 2, and the longest contiguous segment of 0s has length 3.

Note that if there are no 0s, then the longest contiguous segment of 0s is considered to have length 0. The same applies if there are no 1s.

 

Example 1:

```
Input: s = "1101"
Output: true
Explanation:
The longest contiguous segment of 1s has length 2: "1101"
The longest contiguous segment of 0s has length 1: "1101"
The segment of 1s is longer, so return true.
```

Example 2:

```
Input: s = "111000"
Output: false
Explanation:
The longest contiguous segment of 1s has length 3: "111000"
The longest contiguous segment of 0s has length 3: "111000"
The segment of 1s is not longer, so return false.
```

Example 3:

```
Input: s = "110100010"
Output: false
Explanation:
The longest contiguous segment of 1s has length 2: "110100010"
The longest contiguous segment of 0s has length 3: "110100010"
The segment of 1s is not longer, so return false.
```

Constraints:

1 <= s.length <= 100   
s[i] is either '0' or '1'.   


## 方法



```go

func checkZeroOnes(s string) bool {
	cnt0 := 0
	cnt1 := 0
	result0 := 0
	result1 := 0
	for i:=0; i < len(s); i++ {
		if s[i] == '0' {
			cnt0++
			cnt1 = 0
		} else {
			cnt1++
			cnt0 = 0
		}
		result0 = max(result0, cnt0)
		result1 = max(result1, cnt1)
	}
	return result1 > result0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

```


```python
class Solution:
    def checkZeroOnes(self, s: str) -> bool:
        i = 0 
        cnt0 = 0
        cnt1 = 0
        result0 = 0
        result1 = 0
        while i < len(s):
            if s[i] == '0':
                cnt0 += 1
                cnt1 = 0
            if s[i] == '1':
                cnt1 += 1
                cnt0 = 0
            i += 1
            result0 = max(result0, cnt0)
            result1 = max(result1, cnt1)
        return result1 > result0


```