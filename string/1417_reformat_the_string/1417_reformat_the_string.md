1417. Reformat The String


Easy


Given alphanumeric string s. (Alphanumeric string is a string consisting of lowercase English letters and digits).

You have to find a permutation of the string where no letter is followed by another letter and no digit is followed by another digit. That is, no two adjacent characters have the same type.

Return the reformatted string or return an empty string if it is impossible to reformat the string.

 

Example 1:

```
Input: s = "a0b1c2"
Output: "0a1b2c"
Explanation: No two adjacent characters have the same type in "0a1b2c". "a0b1c2", "0a1b2c", "0c2a1b" are also valid permutations.
```

Example 2:

```
Input: s = "leetcode"
Output: ""
Explanation: "leetcode" has only characters so we cannot separate them by digits.
```

Example 3:

```
Input: s = "1229857369"
Output: ""
Explanation: "1229857369" has only digits so we cannot separate them by characters.
```

Example 4:

```
Input: s = "covid2019"
Output: "c2o0v1i9d"
```

Example 5:

```
Input: s = "ab123"
Output: "1a2b3"
```

Constraints:

1 <= s.length <= 500  
s consists of only lowercase English letters and/or digits.

## 方法

```go
func reformat(s string) string {
    cm := make([]byte, 0)
	im := make([]byte, 0)
	for _, v := range []byte(s) {
		if v >= '0' && v <= '9' {
			im = append(im, v)
		} else {
			cm = append(cm, v)
		}
	}

	cmLen := len(cm)
	imLen := len(im)

	if !(cmLen-imLen == 0 || cmLen-imLen == 1 || cmLen-imLen == -1) {
		return ""
	}

	minLen := 0
	flag := 0
	result := ""
	if cmLen == imLen {
		minLen = imLen
	}else if cmLen > imLen {
		minLen = imLen
		flag = 1
	}else{
		minLen = cmLen
		flag = 2
	}

	for i := 0; i < minLen; i++ {
		if flag == 1 {
			result += string(cm[i]) + string(im[i])
		}else{
			result += string(im[i]) + string(cm[i])
		}
	}

	if flag == 1{
		result += string(cm[cmLen-1])
	}else if flag == 2{
		result += string(im[imLen-1])
	}
	return result

}
```



```python
class Solution:
    def reformat(self, s: str) -> str:
        a=re.findall(r'\d',s)
        b=re.findall(r'[a-z]',s)
        if abs(len(a)-len(b))>1:
            return ''
        a,b=sorted([a,b],key=len)
        return ''.join(map(''.join,itertools.zip_longest(b,a,fillvalue='')))
```