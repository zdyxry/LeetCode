167. Two Sum II - Input array is sorted

Easy

Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

Your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution and you may not use the same element twice.
Example:

```
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
```

# 方法
遍历数组，通过字典保存数字，其中 key 为数字，value 为数字索引。  
注意：最终返回的数值应为 `索引 + 1 `




```go
func twoSum(numbers []int, target int) []int {
    
    start, end := 0, len(numbers) - 1
    for start != end {
        sum := numbers[start] + numbers[end]
        if sum > target {
            end -= 1
        } else if sum < target {
            start += 1
        } else {
            break
        }
    }
    return []int{start + 1, end + 1}
}

```


```python
class Solution(object):
    def twoSum(self, numbers, target):
        """
        :type numbers: List[int]
        :type target: int
        :rtype: List[int]
        """
        d = {}
        for idx, num in enumerate(numbers):
            if target - num in d:
                return [d[target-num] + 1 , idx + 1 ]
            d[num] = idx
```