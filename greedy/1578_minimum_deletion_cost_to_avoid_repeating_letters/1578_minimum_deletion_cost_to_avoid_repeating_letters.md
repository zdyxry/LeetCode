1578. Minimum Deletion Cost to Avoid Repeating Letters


Medium


Given a string s and an array of integers cost where cost[i] is the cost of deleting the character i in s.

Return the minimum cost of deletions such that there are no two identical letters next to each other.

Notice that you will delete the chosen characters at the same time, in other words, after deleting a character, the costs of deleting other characters will not change.

 

Example 1:

```
Input: s = "abaac", cost = [1,2,3,4,5]
Output: 3
Explanation: Delete the letter "a" with cost 3 to get "abac" (String without two identical letters next to each other).
```

Example 2:

```
Input: s = "abc", cost = [1,2,3]
Output: 0
Explanation: You don't need to delete any character because there are no identical letters next to each other.
```

Example 3:

```
Input: s = "aabaa", cost = [1,2,3,4,1]
Output: 2
Explanation: Delete the first and the last character, getting the string ("aba").
```

Constraints:

s.length == cost.length  
1 <= s.length, cost.length <= 10^5  
1 <= cost[i] <= 10^4  
s contains only lowercase English letters.


## 方法

```go
func minCost(s string, cost []int) int {
    sum := 0
    for i := range cost {
        sum += cost[i]
    }
    res := 0
    for i := 0; i < len(s);  {
        m := cost[i]
        if i < len(s)- 1 && s[i]==s[i+1] {
            j := i
            for j < len(s) && s[j] == s[i] {
                m = max(m, cost[j])
                j++
            }
            i = j
        } else {
            i++
        }
        res += m
    }
    
    return sum - res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

```


```python
class Solution:
    def minCost(self, s: str, cost: List[int]) -> int:
        pre = 0
        res = 0
        for idx in range(1, len(s)):
            if s[idx] == s[pre]:
                res += min(cost[idx], cost[pre])
            if s[idx] != s[pre] or cost[pre] < cost[idx]:
                pre = idx
        return res
```