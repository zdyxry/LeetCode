942. DI String Match


Easy


Given a string S that only contains "I" (increase) or "D" (decrease), let N = S.length.

Return any permutation A of [0, 1, ..., N] such that for all i = 0, ..., N-1:

```
If S[i] == "I", then A[i] < A[i+1]
If S[i] == "D", then A[i] > A[i+1]
```

Example 1:

```
Input: "IDID"
Output: [0,4,1,3,2]
```

Example 2:

```
Input: "III"
Output: [0,1,2,3]
```

Example 3:

```
Input: "DDI"
Output: [3,2,0,1]
```
 

Note:

1 <= S.length <= 10000   
S only contains characters "I" or "D".


## 方法

```go
func diStringMatch(S string) []int {
    var(
        l, r = 0, len(S)
        ans = make([]int, 0, len(S) + 1)
    )
    for _, s := range S{
        switch s{
            case 'I':
                ans = append(ans, l)
                l ++
            default:
                ans = append(ans, r)
                r --
        }
    }
    ans = append(ans, l)

    return ans
}
```


```python
class Solution:
    def diStringMatch(self, S: str) -> List[int]:
        lo, hi = 0, len(S)
        ans = []
        for x in S:
            if x == 'I':
                ans.append(lo)
                lo += 1
            else:
                ans.append(hi)
                hi -= 1

        return ans + [lo]
```