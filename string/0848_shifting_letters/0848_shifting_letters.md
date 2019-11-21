848. Shifting Letters


Medium


We have a string S of lowercase letters, and an integer array shifts.

Call the shift of a letter, the next letter in the alphabet, (wrapping around so that 'z' becomes 'a'). 

For example, shift('a') = 'b', shift('t') = 'u', and shift('z') = 'a'.

Now for each shifts[i] = x, we want to shift the first i+1 letters of S, x times.

Return the final string after all such shifts to S are applied.

Example 1:

```
Input: S = "abc", shifts = [3,5,9]
Output: "rpl"
Explanation: 
We start with "abc".
After shifting the first 1 letters of S by 3, we have "dbc".
After shifting the first 2 letters of S by 5, we have "igc".
After shifting the first 3 letters of S by 9, we have "rpl", the answer.
```

Note:

1 <= S.length = shifts.length <= 20000  
0 <= shifts[i] <= 10 ^ 9  

## 方法

```go
func shiftingLetters(S string, shifts []int) string {
    size := len(S)
	a := s2is(S)
	shift := 0
	for i := size - 1; 0 <= i; i-- {
		shift += shifts[i]
		shift %= 26
		a[i] = (a[i] + shift) % 26
	}
	return is2s(a)
}

func s2is(s string) []int {
	res := make([]int, len(s))
	for i := range s {
		res[i] = int(s[i] - 'a')
	}
	return res
}

func is2s(a []int) string {
	bs := make([]byte, len(a))
	for i, n := range a {
		bs[i] = byte(n + 'a')
	}
	return string(bs)
}
```


```python
class Solution(object):
    def shiftingLetters(self, S, shifts):
        """
        :type S: str
        :type shifts: List[int]
        :rtype: str
        """
        result = []
        times = sum(shifts) % 26
        for i, c in enumerate(S):
            index = ord(c) - ord('a')
            result.append(chr(ord('a') + (index+times) % 26))
            times = (times-shifts[i]) % 26
        return "".join(result)
```