1209. Remove All Adjacent Duplicates in String II


Medium


Given a string s, a k duplicate removal consists of choosing k adjacent and equal letters from s and removing them causing the left and the right side of the deleted substring to concatenate together.

We repeatedly make k duplicate removals on s until we no longer can.

Return the final string after all such duplicate removals have been made.

It is guaranteed that the answer is unique.

 

Example 1:

```
Input: s = "abcd", k = 2
Output: "abcd"
Explanation: There's nothing to delete.
```

Example 2:

```
Input: s = "deeedbbcccbdaa", k = 3
Output: "aa"
Explanation: 
First delete "eee" and "ccc", get "ddbbbdaa"
Then delete "bbb", get "dddaa"
Finally delete "ddd", get "aa"
```

Example 3:

```
Input: s = "pbbcggttciiippooaais", k = 2
Output: "ps"
```
 

Constraints:

1 <= s.length <= 10^5  
2 <= k <= 10^4  
s only contains lower case English letters.

## 方法

```go
func removeDuplicates(s string, k int) string {
    bytes:=[]byte(s)
	counter:=make([]int,len(s))
	for i:=0;i<len(bytes);i++{
		if i==0||bytes[i]!=bytes[i-1]{
			counter[i]=1
		}else{
			counter[i]=counter[i-1]+1
			if counter[i]==k{
				bytes=append(bytes[:i-k+1],bytes[i+1:]...)
				i=i-k
			}
		}
	}
	return string(bytes)
}
```

```python
class Solution:
    def removeDuplicates(self, s: str, k: int) -> str:
        stack = []
        for i in s:
            print(i)
            if len(stack) == 0 or stack[-1][0] != i:
                stack.append([i, 1])
            else:
                stack[-1][1] += 1
                if stack[-1][1] == k:
                    stack.pop()
                    
        ret = ""
        for t in stack:
            ret += t[0] * t[1]
        return ret
```