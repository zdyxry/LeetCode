187. Repeated DNA Sequences


Medium


All DNA is composed of a series of nucleotides abbreviated as A, C, G, and T, for example: "ACGAATTCCG". When studying DNA, it is sometimes useful to identify repeated sequences within the DNA.

Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.

Example:

```
Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"

Output: ["AAAAACCCCC", "CCCCCAAAAA"]
```


## 方法



```go
func findRepeatedDnaSequences(s string) []string {
    var res []string
	if len(s) <= 10 {
		return nil
	}

	str := ""
	// rec 记录各个 subString 出现的次数
	rec := make(map[string]int, len(s)-9)
	for i := 0; i+10 <= len(s); i++ {
		str = s[i : i+10]
		if v := rec[str]; v == 1 {
			// 把已经出现过一次的 subString 存放入 res
			res = append(res, str)
		}
		rec[str]++
	}

	sort.Strings(res)

	return res
}
```


```python
class Solution(object):
    def findRepeatedDnaSequences(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        dict, rolling_hash, res = {}, 0, []

        for i in xrange(len(s)):
            rolling_hash = ((rolling_hash << 3) & 0x3fffffff) | (ord(s[i]) & 7)
            if rolling_hash not in dict:
                dict[rolling_hash] = True
            elif dict[rolling_hash]:
                res.append(s[i - 9: i + 1])
                dict[rolling_hash] = False
        return res
```