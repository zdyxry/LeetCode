830. Positions of Large Groups

Easy

In a string S of lowercase letters, these letters form consecutive groups of the same character.

For example, a string like S = "abbxxxxzyy" has the groups "a", "bb", "xxxx", "z" and "yy".

Call a group large if it has 3 or more characters.  We would like the starting and ending positions of every large group.

The final answer should be in lexicographic order.

 

Example 1:

```
Input: "abbxxxxzzy"
Output: [[3,6]]
Explanation: "xxxx" is the single large group with starting  3 and ending positions 6.
```

Example 2:

```
Input: "abc"
Output: []
Explanation: We have "a","b" and "c" but no large group.
```

Example 3:

```
Input: "abcdddeeeeaabbbcd"
Output: [[3,5],[6,9],[12,14]]
```
 
Note:  1 <= S.length <= 1000


# 方法
暴力求解，当长度大于 3 时，保存。


```go
func largeGroupPositions(S string) [][]int {
    
    res := make([][]int, 0, len(S)/3)
	l, r := 0, 1

	for ; r < len(S); r++ {
		if S[l] != S[r] {
			l = r
			continue
		}

		if r-l+1 == 3 { 
			res = append(res, []int{l, r})
		} else if r-l+1 > 3 { 
			res[len(res)-1][1] = r
		}
	}

	return res
    
}
```


```python
class Solution(object):
    def largeGroupPositions(self, S):
        """
        :type S: str
        :rtype: List[List[int]]
        """
        result = []
        i = 0
        for j in xrange(len(S)):
            if j == len(S)-1 or S[j] != S[j+1]:
                if j-i+1 >= 3:
                    result.append([i, j])
                i = j+1
        return result

```