946. Validate Stack Sequences


Medium


Given two sequences pushed and popped with distinct values, return true if and only if this could have been the result of a sequence of push and pop operations on an initially empty stack.

 
Example 1:

```
Input: pushed = [1,2,3,4,5], popped = [4,5,3,2,1]
Output: true
Explanation: We might do the following sequence:
push(1), push(2), push(3), push(4), pop() -> 4,
push(5), pop() -> 5, pop() -> 3, pop() -> 2, pop() -> 1
```

Example 2:

```
Input: pushed = [1,2,3,4,5], popped = [4,3,5,1,2]
Output: false
Explanation: 1 cannot be popped before 2.
```
 

Note:

0 <= pushed.length == popped.length <= 1000   
0 <= pushed[i], popped[i] < 1000   
pushed is a permutation of popped.   
pushed and popped have distinct values.   


## 方法


```go
func validateStackSequences(pushed []int, popped []int) bool {
    size := len(pushed)

	stack := make([]int, size)
	top := -1

	for in, out := 0, 0; in < size; in++ {
		if pushed[in] != popped[out] {
			top++
			stack[top] = pushed[in]
		} else {
			out++
			for top >= 0 && stack[top] == popped[out] {
				top--
				out++
			}
		}
	}

	return top == -1
}
```


```python
class Solution(object):
    def validateStackSequences(self, pushed, popped):
        """
        :type pushed: List[int]
        :type popped: List[int]
        :rtype: bool
        """
        i = 0
        s = []
        for v in pushed:
            s.append(v)
            while s and i < len(popped) and s[-1] == popped[i]:
                s.pop()
                i += 1
        return i == len(popped)
```