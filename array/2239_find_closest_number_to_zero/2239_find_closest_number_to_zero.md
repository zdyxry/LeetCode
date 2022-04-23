2239. Find Closest Number to Zero


Easy


Given an integer array nums of size n, return the number with the value closest to 0 in nums. If there are multiple answers, return the number with the largest value.

 

Example 1:

```
Input: nums = [-4,-2,1,4,8]
Output: 1
Explanation:
The distance from -4 to 0 is |-4| = 4.
The distance from -2 to 0 is |-2| = 2.
The distance from 1 to 0 is |1| = 1.
The distance from 4 to 0 is |4| = 4.
The distance from 8 to 0 is |8| = 8.
Thus, the closest number to 0 in the array is 1.
```

Example 2:

```
Input: nums = [2,-1,1]
Output: 1
Explanation: 1 and -1 are both the closest numbers to 0, so 1 being larger is returned.
```

Constraints:

1 <= n <= 1000  
-105 <= nums[i] <= 105   


## 方法


```
func findClosestNumber(nums []int) int {
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		if abs(nums[i]) < abs(res) || abs(nums[i]) == abs(res) && nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

```


```
class Solution:
    def findClosestNumber(self, nums: List[int]) -> int:
        result = float('inf')
        min_value = float('inf')
        for n in nums:
            if abs(n-0) == min_value:
                result = max(result, n)
            if abs(n-0) < min_value:
                result = n
                min_value = abs(n-0)
        return result
```