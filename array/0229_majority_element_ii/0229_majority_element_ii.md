229. Majority Element II


Medium


Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

Note: The algorithm should run in linear time and in O(1) space.

Example 1:
```
Input: [3,2,3]
Output: [3]
```

Example 2:

```
Input: [1,1,1,3,3,2,2,2]
Output: [1,2]
```


## 方法

数组中出现次数超过 n/3 次的数字最多有两个。

```go
func majorityElement(nums []int) []int {
    count1, count2, candidate1, candidate2 := 0, 0, 0, 1
	// Select Candidates
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 <= 0 {
			// We have a bad first candidate, replace!
			candidate1, count1 = num, 1
		} else if count2 <= 0 {
			// We have a bad second candidate, replace!
			candidate2, count2 = num, 1
		} else {
			// Both candidates suck, boo!
			count1--
			count2--
		}
	}
	// Recount!
	count1, count2 = 0, 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}
	length := len(nums)
	if count1 > length/3 && count2 > length/3 {
		return []int{candidate1, candidate2}
	}
	if count1 > length/3 {
		return []int{candidate1}
	}
	if count2 > length/3 {
		return []int{candidate2}
	}
	return []int{}
}
```


```python
class Solution(object):
    def majorityElement(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        if(not nums):
            return []
        
        num1 = -1
        num2 = -1
        count1 = 0
        count2 = 0
        
        for n in nums:
            if(n == num1):
                count1 = count1+1
            elif(n == num2):
                count2 = count2+1
            elif(count1 == 0):
                num1 = n
                count1 = 1
            elif(count2 == 0):
                num2 = n
                count2 = 1
            else:
                count1 = count1-1
                count2 = count2-1
                
        return[n for n in [num1, num2] if nums.count(n) > len(nums) // 3]
```