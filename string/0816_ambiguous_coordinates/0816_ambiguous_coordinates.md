816. Ambiguous Coordinates


Medium


We had some 2-dimensional coordinates, like "(1, 3)" or "(2, 0.5)".  Then, we removed all commas, decimal points, and spaces, and ended up with the string S.  Return a list of strings representing all possibilities for what our original coordinates could have been.

Our original representation never had extraneous zeroes, so we never started with numbers like "00", "0.0", "0.00", "1.0", "001", "00.01", or any other number that can be represented with less digits.  Also, a decimal point within a number never occurs without at least one digit occuring before it, so we never started with numbers like ".1".

The final answer list can be returned in any order.  Also note that all coordinates in the final answer have exactly one space between them (occurring after the comma.)

Example 1:
```
Input: "(123)"
Output: ["(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"]
```

Example 2:
```
Input: "(00011)"
Output:  ["(0.001, 1)", "(0, 0.011)"]
Explanation: 
0.0, 00, 0001 or 00.01 are not allowed.
```

Example 3:
```
Input: "(0123)"
Output: ["(0, 123)", "(0, 12.3)", "(0, 1.23)", "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)"]
```

Example 4:
```
Input: "(100)"
Output: [(10, 0)]
Explanation: 
1.0 is not allowed.
```

 

Note:

4 <= S.length <= 12.  
S[0] = "(", S[S.length - 1] = ")", and the other elements in S are digits.  

## 方法


```go
func ambiguousCoordinates(S string) []string {
	res := make([]string, 0, len(S))
	s := S[1 : len(S)-1]
	for i := 1; i < len(s); i++ {
		lefts, rights := addDot(s[:i]), addDot(s[i:])
		for _, l := range lefts {
			for _, r := range rights {
				res = append(res, connect(l, r))
			}
		}
	}
	return res
}

func addDot(s string) []string {
	res := make([]string, 0, len(s))
	if isValid(s) {
		res = append(res, s)
	}

	for i := 1; i < len(s); i++ {
		t := s[:i] + "." + s[i:]
		if isValid(t) {
			res = append(res, t)
		}
	}

	return res
}

func isValid(s string) bool {
	f, _ := strconv.ParseFloat(s, 64)
	fs := strconv.FormatFloat(f, 'f', -1, 64)
	return s == fs
}

func connect(left, right string) string {
	return fmt.Sprintf("(%s, %s)", left, right)
}
```


```python
class Solution(object):
    def ambiguousCoordinates(self, S):
        """
        :type S: str
        :rtype: List[str]
        """
        S = S[1:-1]
        res = []
            # 加小数点
        def f(S):
            if not S or (len(S) > 1 and S[0] == "0" and S[-1] == "0"):
                return []
            if S == "0":
                return [S[0]]
            if S[0] == "0" and len(S) > 1:
                return [S[0]+ "." + S[1:]]
            if S[-1] =="0" and len(S)  > 1:
                return [S]
            return [S] +  [S[:i] + "." + S[i:] for i in range(1, len(S))]
            # 拆数字
        for i in range(1,len(S)):
            for x,y in itertools.product(f(S[:i]), f(S[i:])):
                res.append("(%s, %s)"%(x,y))
        return res
```