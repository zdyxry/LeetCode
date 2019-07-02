344. Reverse String


Easy


Write a function that reverses a string. The input string is given as an array of characters char[].

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

You may assume all the characters consist of printable ascii characters.

 

Example 1:
```
Input: ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
```

Example 2:

```
Input: ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
```


## 方法
双指针，随着索引增加，交换位置。

python 可以直接用内置方法求解。

（这题的运行时间真是有问题。。




```go
func reverseString(s []byte)  {
    start := 0 
    end := len(s) - 1
    for ; start < end ; {
        s[start], s[end] = s[end], s[start]
        start++
        end--
    }
}
```


```python
class Solution(object):
    def reverseString(self, s):
        """
        :type s: str
        :rtype: str
        """
        s[:]=s[::-1]

```