1576. Replace All ?'s to Avoid Consecutive Repeating Characters


Easy


Given a string s containing only lower case English letters and the '?' character, convert all the '?' characters into lower case letters such that the final string does not contain any consecutive repeating characters. You cannot modify the non '?' characters.

It is guaranteed that there are no consecutive repeating characters in the given string except for '?'.

Return the final string after all the conversions (possibly zero) have been made. If there is more than one solution, return any of them. It can be shown that an answer is always possible with the given constraints.

 

Example 1:

```
Input: s = "?zs"
Output: "azs"
Explanation: There are 25 solutions for this problem. From "azs" to "yzs", all are valid. Only "z" is an invalid modification as the string will consist of consecutive repeating characters in "zzs".
```

Example 2:

```
Input: s = "ubv?w"
Output: "ubvaw"
Explanation: There are 24 solutions for this problem. Only "v" and "w" are invalid modifications as the strings will consist of consecutive repeating characters in "ubvvw" and "ubvww".
```

Example 3:

```
Input: s = "j?qg??b"
Output: "jaqgacb"
```

Example 4:

```
Input: s = "??yw?ipkj?"
Output: "acywaipkja"
```

Constraints:

1 <= s.length <= 100

s contains only lower case English letters and '?'.


## 方法



```go
func modifyString(s string) string {
    chars := make([]byte, len(s))
	for i := 0; i < len(chars); i++ {
		if s[i] != '?' {
			chars[i] = s[i]
			continue
		}
		for j := byte('a'); j <= 'z'; j++ {
			if (i == 0 || j != chars[i-1]) && (i == len(chars)-1 || j != s[i+1]) {
				chars[i] = j
				break
			}
		}
	}
	return string(chars)
}
```

```python
class Solution:
    def modifyString(self, s: str) -> str:
        l = list(s)
        alphabet = list(string.ascii_lowercase)
        for i, _ in enumerate(l):
            if l[i] == "?":
                for j in alphabet:
                    if i == len(l) - 1:
                        if l[i-1] != j:
                            l[i] = j
                            break
                    elif l[i-1] != j and l[i+1] != j:
                        l[i] = j
                        break
        return ''.join(l)
```