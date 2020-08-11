1546. Maximum Number of Non-Overlapping Subarrays With Sum Equals Target


Medium


Given an array nums and an integer target.

Return the maximum number of non-empty non-overlapping subarrays such that the sum of values in each subarray is equal to target.

 

Example 1:

```
Input: nums = [1,1,1,1,1], target = 2
Output: 2
Explanation: There are 2 non-overlapping subarrays [1,1,1,1,1] with sum equals to target(2).
```

Example 2:

```
Input: nums = [-1,3,5,1,4,2,-9], target = 6
Output: 2
Explanation: There are 3 subarrays with sum equal to 6.
([5,1], [4,2], [3,5,1,4,2,-9]) but only the first 2 are non-overlapping.
```

Example 3:

```
Input: nums = [-2,6,6,3,5,4,1,2,8], target = 10
Output: 3
```

Example 4:

```
Input: nums = [0,0,0], target = 0
Output: 3
```
 

Constraints:

1 <= nums.length <= 10^5  
-10^4 <= nums[i] <= 10^4  
0 <= target <= 10^6


## 方法

```go
func maxNonOverlapping(nums []int, target int) int {
    count,sum := 0,0
	m := make(map[int]int)
	m[0]=1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		_,hasKey:=m[sum-target];
		if  hasKey{
			m = make(map[int]int)
			count++
			sum = 0
		}
		m[sum] ++
	}

	return count
}
```

```python
class Solution:
    def maxNonOverlapping(self, nums: List[int], target: int) -> int:
        s = {0}
        cur_sum = 0
        ans = 0
        for num in nums:
            cur_sum += num
            if cur_sum - target in s:
                s = {0}
                cur_sum = 0
                ans += 1
            else:
                s.add(cur_sum)
        return ans
```