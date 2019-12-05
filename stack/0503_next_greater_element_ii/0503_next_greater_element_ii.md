503. Next Greater Element II


Medium


Given a circular array (the next element of the last element is the first element of the array), print the Next Greater Number for every element. The Next Greater Number of a number x is the first greater number to its traversing-order next in the array, which means you could search circularly to find its next greater number. If it doesn't exist, output -1 for this number.

Example 1:
```
Input: [1,2,1]
Output: [2,-1,2]
Explanation: The first 1's next greater number is 2; 
The number 2 can't find next greater number; 
The second 1's next greater number needs to search circularly, which is also 2.
```

Note: The length of given array won't exceed 10000.


## 方法

```go
func nextGreaterElements(nums []int) []int {
    result := make([]int, len(nums))
    stack := make([]int, len(nums))
    top := -1
    for i := range result {
        result[i] = -1
    }
    for i := 0; i < len(nums) * 2; i++ {
        num := nums[i % len(nums)]
        for top > -1 && nums[stack[top]] < num {
            result[stack[top]] = num
            top--
        }
        if i < len(nums) {
            top++
            stack[top] = i
        }
    }

    return result
}
```


```python
class Solution(object):
    def nextGreaterElements(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        stack = []
        result = [-1]*len(nums)
        for enum, n in enumerate(nums):
            while stack and stack[-1][0] < n:
                result[stack[-1][1]] = n
                stack.pop()
            stack.append((n, enum))
            
        if stack:
            for enum, n in enumerate(nums):
                while stack and stack[-1][0] < n:
                    result[stack[-1][1]] = n
                    stack.pop()
                    
                
        return result
```