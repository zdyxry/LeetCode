561. Array Partition I

Easy

Given an array of 2n integers, your task is to group these integers into n pairs of integer, say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of min(ai, bi) for all i from 1 to n as large as possible.

Example 1:
```
Input: [1,4,3,2]

Output: 4
Explanation: n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).
```


Note:
n is a positive integer, which is in the range of [1, 10000].
All the integers in the array will be in the range of [-10000, 10000].


# 方法
拆分为两个数字一组之后，每对数字之差应尽可能的小，保证不会浪费过多的间隔。  
1. 先排序，然后步长为 2 进行相加。
2. 题目给定数字范围是 [-10000, 10000]，那么将数组进行偏移， `索引 - 10000` 即为数字本身，索引为 `索引 - 10000` 的值为该数字出现的次数，每次只对第一次出现数字相加。



```go
func arrayPairSum(nums []int) int {
    var buckets [20001]int8
	for _, num := range nums {
		buckets[num+10000]++
	}

	sum := 0
	needAdd := true
	for num, count := range buckets {
        
		for count > 0 {
			if needAdd {
				sum += num - 10000
			}
			needAdd = !needAdd
			count--
		}
	}

	return sum
}
```


```python
class Solution(object):
    def arrayPairSum(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        nums.sort()
        result = 0
        for i in xrange(0, len(nums), 2):
            result += nums[i]
        return result
```