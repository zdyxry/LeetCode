784. Letter Case Permutation


Medium

Given a string S, we can transform every letter individually to be lowercase or uppercase to create another string.

Return a list of all possible strings we could create. You can return the output in any order.

 

Example 1:

```
Input: S = "a1b2"
Output: ["a1b2","a1B2","A1b2","A1B2"]
```

Example 2:

```
Input: S = "3z4"
Output: ["3z4","3Z4"]
```

Example 3:

```
Input: S = "12345"
Output: ["12345"]
```

Example 4:

```
Input: S = "0"
Output: ["0"]
```
 

Constraints:

S will be a string with length between 1 and 12.   
S will consist only of letters or digits.


## 方法


```go
func letterCasePermutation(S string) []string {
    permutations := []string{S}
	for i, v := range S {
		if v >= 'a' && v <= 'z' {
			length := len(permutations)
			for ip := 0; ip < length; ip++ {
				b := []byte(permutations[ip])
				b[i] = byte(v +'A'-'a')
				permutations = append(permutations, string(b))
			}
		} else if v >= 'A' && v <= 'Z' {
			length := len(permutations)
			for ip := 0; ip < length; ip++ {
				b := []byte(permutations[ip])
				b[i] = byte(v +'a'-'A')
				permutations = append(permutations, string(b))
			}
		}
	}
	return permutations
}
```


```python
class Solution:
    def letterCasePermutation(self, S: str) -> List[str]:
        ans = [[]]

        for char in S:
            n = len(ans)
            if char.isalpha():
                for i in range(n):
                    ans.append(ans[i][:])
                    ans[i].append(char.lower())
                    ans[n+i].append(char.upper())
                    print(ans)
            else:
                for i in range(n):
                    ans[i].append(char)

        return map("".join, ans)
```