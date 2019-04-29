645. Set Mismatch
Easy

The set S originally contains numbers from 1 to n. But unfortunately, due to the data error, one of the numbers in the set got duplicated to another number in the set, which results in repetition of one number and loss of another number.

Given an array nums representing the data status of this set after the error. Your task is to firstly find the number occurs twice and then find the number that is missing. Return them in the form of an array.

Example 1:  
Input: nums = [1,2,2,4]  
Output: [2,3]  
Note:  
The given array size will in the range [2, 10000].  
The given array's numbers won't have any order.  

# 方法
1. 数学方法，根据 nums 长度，算出正确长度所有数字之和，减去去重后 nums 之和，为缺失数字    

2. 遍历数组，当 nums[i] - 1 为正时，则将其标记为负，直到遇到重复数字


```go

func findErrorNums(nums []int) []int {
	dup := 0
	for i := 0; i < len(nums); i++ {
		n := abs(nums[i])

		if nums[n-1] < 0 {
			dup = n
		} else {
			nums[n-1] = -nums[n-1]
		}
	}

	mis := 0
	for i, v := range nums {
		if v > 0 {
			mis = i + 1
			break
		}
	}

	return []int{dup, mis}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
```


```python
def findErrorNums(self, nums):
    return [sum(nums) - sum(set(nums)), sum(range(1, len(nums)+1)) - sum(set(nums))]
```

```python
class Solution(object):
    def findErrorNums(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        result = [0] * 2
        for i in nums:
            if nums[abs(i)-1] < 0:
                result[0] = abs(i)
            else:
                nums[abs(i)-1] *= -1
        for i in xrange(len(nums)):
            if nums[i] > 0:
                result[1] = i+1
            else:
                nums[i] *= -1
        return result
```