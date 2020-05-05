1433. Check If a String Can Break Another String


Medium


Given two strings: s1 and s2 with the same size, check if some permutation of string s1 can break some permutation of string s2 or vice-versa (in other words s2 can break s1).

A string x can break string y (both of size n) if x[i] >= y[i] (in alphabetical order) for all i between 0 and n-1.

 

Example 1:

```
Input: s1 = "abc", s2 = "xya"
Output: true
Explanation: "ayx" is a permutation of s2="xya" which can break to string "abc" which is a permutation of s1="abc".
```

Example 2:

```
Input: s1 = "abe", s2 = "acd"
Output: false 
Explanation: All permutations for s1="abe" are: "abe", "aeb", "bae", "bea", "eab" and "eba" and all permutation for s2="acd" are: "acd", "adc", "cad", "cda", "dac" and "dca". However, there is not any permutation from s1 which can break some permutation from s2 and vice-versa.
```

Example 3:

```
Input: s1 = "leetcodee", s2 = "interview"
Output: true
```
 

Constraints:

s1.length == n  
s2.length == n  
1 <= n <= 10^5  
All strings consist of lowercase English letters.

## 方法

```go
func checkIfCanBreak(s1 string, s2 string) bool {
    return helper(s1, s2) || helper(s2, s1)
}

func helper(s1 string, s2 string) bool {
    m2 := map[int]int{}
    for i := 0; i < len(s2); i++ {
        m2[int(s2[i] - 'a')]++
    }
    for i := 0; i < len(s1); i++ {
        j := int(s1[i] - 'a')
        for ; j <= 26; j++ {
            if m2[j] > 0 {
                m2[j]--
                break
            }
        }
        if j > 26 {
            return false
        }
    }
    return true
}
```


```python
class Solution:
    def checkIfCanBreak(self, s1: str, s2: str) -> bool:
        s1= sorted(s1)
        s2= sorted(s2)
        return all(s1[i] >= s2[i] for i in range(len(s1))) or all(s1[i] <= s2[i] for i in range(len(s1)))
```