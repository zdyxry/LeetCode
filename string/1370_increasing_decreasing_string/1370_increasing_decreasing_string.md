1370. Increasing Decreasing String


Easy


Given a string s. You should re-order the string using the following algorithm:

Pick the smallest character from s and append it to the result.

Pick the smallest character from s which is greater than the last appended character to the result and append it.

Repeat step 2 until you cannot pick more characters.

Pick the largest character from s and append it to the result.

Pick the largest character from s which is smaller than the last appended character to the result and append it.

Repeat step 5 until you cannot pick more characters.

Repeat the steps from 1 to 6 until you pick all characters from s.

In each step, If the smallest or the largest character appears more than once you can choose any occurrence and append it to the result.


Return the result string after sorting s with this algorithm.

 

Example 1:

```
Input: s = "aaaabbbbcccc"
Output: "abccbaabccba"
Explanation: After steps 1, 2 and 3 of the first iteration, result = "abc"
After steps 4, 5 and 6 of the first iteration, result = "abccba"
First iteration is done. Now s = "aabbcc" and we go back to step 1
After steps 1, 2 and 3 of the second iteration, result = "abccbaabc"
After steps 4, 5 and 6 of the second iteration, result = "abccbaabccba"
```

Example 2:

```
Input: s = "rat"
Output: "art"
Explanation: The word "rat" becomes "art" after re-ordering it with the mentioned algorithm.
```

Example 3:

```
Input: s = "leetcode"
Output: "cdelotee"
```

Example 4:

```
Input: s = "ggggggg"
Output: "ggggggg"
```

Example 5:

```
Input: s = "spo"
Output: "ops"
```
 

Constraints:

1 <= s.length <= 500

s contains only lower-case English letters.

## 方法


```go
func sortString(s string) string {
	sMap := make(map[string]int, len(s))
	for _, c := range s {
		sMap[string(c)]++
	}
	chars := []string{}
	for c, _ := range sMap {
		chars = append(chars, string(c))
	}
	sort.Strings(chars)
	res := ""
	for len(sMap) > 0 {
		for i := range chars {
			if sMap[chars[i]] > 0 {
				res += string(chars[i])
				sMap[chars[i]]--
			} else {
				delete(sMap, chars[i])
			}
		}
		for i := range chars {
			if sMap[chars[len(chars)-i-1]] > 0 {
				res += string(chars[len(chars)-i-1])
				sMap[chars[len(chars)-i-1]]--
			} else {
				delete(sMap, chars[len(chars)-i-1])
			}
		}
	}
	return res
}
```




```python
class Solution(object):
    def sortString(self, s):
        """
        :type s: str
        :rtype: str
        """
        set_s = set()
        for c in s:
			set_s.add(c)
        characters = sorted(set_s)

        d = collections.Counter(s)
        res = ""
        while sum(d.values()) != 0:
            for c in characters:
                if c in s and d[c] > 0:
                    res += c
                    d[c] -= 1
            for c in characters[::-1]:
                if c in s and d[c] > 0:
                    res += c
                    d[c] -= 1
        return res
```