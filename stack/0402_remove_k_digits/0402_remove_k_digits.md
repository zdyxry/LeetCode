402. Remove K Digits


Medium


Given a non-negative integer num represented as a string, remove k digits from the number so that the new number is the smallest possible.

Note:  
The length of num is less than 10002 and will be ≥ k.  
The given num does not contain any leading zero.  
Example 1:

```
Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
```

Example 2:

```
Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.
```

Example 3:

```
Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.
```

## 方法


```go
func removeKdigits(num string, k int) string {
    // 返回值的长度
	digits := len(num) - k
	stack := make([]byte, len(num))
	top := 0

	for i := range num {
		// 在还能删除的前提下
		// 从上往下，删除 stack 中所有比 num[i] 大的数
		for top > 0 && stack[top-1] > num[i] && k > 0 {
			top--
			k--
		}
		stack[top] = num[i]
		top++
	}

	// 处理开头的　'0'
	i := 0
	for i < digits && stack[i] == '0' {
		i++
	}

	if i == digits {
		return "0"
	}
	return string(stack[i:digits])
}
```


```python
class Solution(object):
    def removeKdigits(self, num, k):
        """
        :type num: str
        :type k: int
        :rtype: str
        compare the largest and second largest significant digit while k>0
        Greedy algorithm
        
        1432219  1 4 remove 4 -> 132219
        132219 1 3 remove 3 -> 12219
        12219 1 2 remove 2 -> 1219
        
        10200 1 0 remove 1 -> 0200 -> 200 
        if k ==2 
        200 2 0 remove 2 -> 00 -> 0
        
        recast to integer to remove leading zeros
        
        Greedy will not work ex. 112  1 1 remove 1 -> 12
        112 remove 2 -> 11
        1432219 132219, 143221, 143219 142219, 432219,
        
        cannot convert to integer
        
        221 -> 21 
        """
            
        if len(num) <= k:
            return "0"
        
        ret = []
        for i in range(len(num)):
            while k > 0 and ret and ret[-1] > num[i]:
                ret.pop()
                k-=1
            ret.append(num[i])
        while k>0 and ret:
            ret.pop()
            k-=1
            
            
        s =  "".join(ret).lstrip("0")
        return s if s else "0"

```