1309. Decrypt String from Alphabet to Integer Mapping


Easy


Given a string s formed by digits ('0' - '9') and '#' . We want to map s to English lowercase characters as follows:

Characters ('a' to 'i') are represented by ('1' to '9') respectively.

Characters ('j' to 'z') are represented by ('10#' to '26#') respectively. 

Return the string formed after mapping.

It's guaranteed that a unique mapping will always exist.

 

Example 1:

```
Input: s = "10#11#12"
Output: "jkab"
Explanation: "j" -> "10#" , "k" -> "11#" , "a" -> "1" , "b" -> "2".
```

Example 2:

```
Input: s = "1326#"
Output: "acz"
```

Example 3:

```
Input: s = "25#"
Output: "y"
```

Example 4:

```
Input: s = "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"
Output: "abcdefghijklmnopqrstuvwxyz"
```
 

Constraints:

1 <= s.length <= 1000  
s[i] only contains digits letters ('0'-'9') and '#' letter.  
s will be valid string such that mapping is always possible.

## 方法

```go
func freqAlphabets(s string) string {
    rs := []byte{}
	for i := 0; i < len(s); i++ {
		if i < len(s) - 2 && s[i+2] == '#' {
			if s[i] == '1' {
				rs = append(rs, s[i+1]-'0'+'j')
			} else {
				rs = append(rs, s[i+1]-'0'+'t')
			}
			i += 2
		} else {
			rs = append(rs, s[i]-'1'+'a')
		}
	}
	return string(rs)
}
```



```python
class Solution(object):
    def freqAlphabets(self, s):
        """
        :type s: str
        :rtype: str
        """
        for i in range(10,27):
            s = s.replace( str(i) + '#',chr(ord('j') - 10 + i) )
        for i in range(1,10):
            s = s.replace( str(i), chr(ord('a') - 1 + i ) )
        return s
```



```python
class Solution(object):
    def freqAlphabets(self, s):
        """
        :type s: str
        :rtype: str
        """
        def get(st):
            print st
            return chr(int(st) + 96)
        
        i, res = 0, ""
        while i < len(s):
            if i + 2 < len(s) and s[i+2] == "#":
                res += get(s[i : i + 2])
                i += 2
            else:
                res += get(s[i])
            i += 1
        return res
```