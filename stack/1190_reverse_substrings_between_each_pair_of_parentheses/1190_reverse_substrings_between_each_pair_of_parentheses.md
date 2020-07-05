1190. Reverse Substrings Between Each Pair of Parentheses


Medium

You are given a string s that consists of lower case English letters and brackets. 

Reverse the strings in each pair of matching parentheses, starting from the innermost one.

Your result should not contain any brackets.

 

Example 1:

```
Input: s = "(abcd)"
Output: "dcba"
```

Example 2:

```
Input: s = "(u(love)i)"
Output: "iloveu"
Explanation: The substring "love" is reversed first, then the whole string is reversed.
```

Example 3:

```
Input: s = "(ed(et(oc))el)"
Output: "leetcode"
Explanation: First, we reverse the substring "oc", then "etco", and finally, the whole string.
```

Example 4:

```
Input: s = "a(bcdefghijkl(mno)p)q"
Output: "apmnolkjihgfedcbq"
```
 

Constraints:

0 <= s.length <= 2000  
s only contains lower case English characters and parentheses.  
It's guaranteed that all parentheses are balanced.


## 方法

```go
func reverseParentheses(s string) string {
    var f []int;
    var chs []byte;
    for i:=0; i<len(s); i++{
        m := s[i]
        if m=='(' {
            f = append(f,len(chs))
        }else if m==')'{
            l,r := f[len(f)-1],len(chs)-1
            for l<r {
                chs[l],chs[r] = chs[r],chs[l]
                l+=1;
                r-=1;
            }
            f = f[:len(f)-1]
        }else {
            chs = append(chs,m)
        }
    }
    return string(chs)
}
```


```python
class Solution:
    def reverseParentheses(self, s: str) -> str:
        stack = ['']
        for c in s:
            if c == '(':
                stack.append('')
            elif c == ')':
                reversed_top = stack.pop()[::-1]
                stack[-1] += reversed_top
            else:
                stack[-1] += c
            print(stack)
        return stack[0]
```