202. Happy Number

Easy

Write an algorithm to determine if a number is "happy".

A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

Example: 

```
Input: 19
Output: true
Explanation: 
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
```



# 方法
通过字典记录出现过的数字，若数字未出现过且唯一，则为真。


```go
func isHappy(n int) bool {
	slow, fast := n, trans(n)
	for slow != fast {
		slow = trans(slow)
		fast = trans(trans(fast))
	}
	if slow == 1 {
		return true
	}
	return false
}

func trans(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}
```


```python
class Solution(object):
    def isHappy(self, n):
        lookup = {}
        while n != 1 and n not in lookup:
            lookup[n] = True
            n = self.nextNumber(n)
        return n == 1

    def nextNumber(self, n):
        new = 0
        for char in str(n):
            new += int(char)**2
        return new
```