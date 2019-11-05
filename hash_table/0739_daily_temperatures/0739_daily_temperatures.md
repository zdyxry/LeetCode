739. Daily Temperatures


Medium


Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature. If there is no future day for which this is possible, put 0 instead.

For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].

Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100].

## 方法

```go
func dailyTemperatures(T []int) []int {
    res := make([]int, len(T), len(T))
    stack := make([]int, 1, len(T))
    
    for i := 1; i < len(T); i++ {
        for len(stack) > 0 && T[i] > T[stack[len(stack) - 1]] {
            idx := stack[len(stack) - 1]
            res[idx] = i - idx
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, i)
    }
    
    return res
}
```


```python
class Solution(object):
    def dailyTemperatures(self, T):
        """
        :type T: List[int]
        :rtype: List[int]
        """
        result = [0] * len(T)
        stk = []
        for i in xrange(len(T)):
            while stk and \
                  T[stk[-1]] < T[i]:
                idx = stk.pop()
                result[idx] = i-idx
            stk.append(i)
        return result
```