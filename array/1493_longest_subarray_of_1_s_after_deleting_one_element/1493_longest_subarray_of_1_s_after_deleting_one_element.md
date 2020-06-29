1493. Longest Subarray of 1's After Deleting One Element


Medium


Given a binary array nums, you should delete one element from it.

Return the size of the longest non-empty subarray containing only 1's in the resulting array.

Return 0 if there is no such subarray.

 

Example 1:

```
Input: nums = [1,1,0,1]
Output: 3
Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's.
```

Example 2:

```
Input: nums = [0,1,1,1,0,1,1,0,1]
Output: 5
Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest subarray with value of 1's is [1,1,1,1,1].
```

Example 3:

```
Input: nums = [1,1,1]
Output: 2
Explanation: You must delete one element.
```

Example 4:

```
Input: nums = [1,1,0,0,1,1,1,0,1]
Output: 4
```

Example 5:

```
Input: nums = [0,0,0]
Output: 0
```
 

Constraints:

1 <= nums.length <= 10^5  
nums[i] is either 0 or 1.

## 方法

```go
func longestSubarray(nums []int) int {
    l, r := 0, 0
    zeroIdx := -1
    
    maxOnes := 0
    for r < len(nums) {
        if nums[r] == 0 {
            l = zeroIdx + 1
            zeroIdx = r
        }
        maxOnes = getMax(maxOnes, r - l)
        r++
    }
    
    return maxOnes
}

func getMax(i, j int) int {
    if i > j {
        return i
    }
    return j
}
```


```python
class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        cnt1, cnt2 = 0, 0
        maxlen = 0
        for num in nums:
            if num == 1:
                cnt1 += 1
                cnt2 += 1
                maxlen = max(maxlen, cnt1)
            else:
                cnt1 = cnt2
                cnt2 = 0
        if cnt1 == len(nums):
            return len(nums)-1
        return maxlen
```


```python
class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        if 0 not in nums:
            return len(nums)-1
        s = ''.join(list(map(str,nums))).split('0')
        res = 0
        for i in range(len(s)-1):
            res = max(res,len(s[i])+len(s[i+1]))
        return res
```
