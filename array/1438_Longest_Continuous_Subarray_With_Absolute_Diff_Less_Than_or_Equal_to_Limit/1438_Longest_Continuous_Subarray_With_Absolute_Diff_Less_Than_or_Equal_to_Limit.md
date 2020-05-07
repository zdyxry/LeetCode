1438. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit


Medium


Given an array of integers nums and an integer limit, return the size of the longest continuous subarray such that the absolute difference between any two elements is less than or equal to limit.

In case there is no subarray satisfying the given condition return 0.

 

Example 1:

```
Input: nums = [8,2,4,7], limit = 4
Output: 2 
Explanation: All subarrays are: 
[8] with maximum absolute diff |8-8| = 0 <= 4.
[8,2] with maximum absolute diff |8-2| = 6 > 4. 
[8,2,4] with maximum absolute diff |8-2| = 6 > 4.
[8,2,4,7] with maximum absolute diff |8-2| = 6 > 4.
[2] with maximum absolute diff |2-2| = 0 <= 4.
[2,4] with maximum absolute diff |2-4| = 2 <= 4.
[2,4,7] with maximum absolute diff |2-7| = 5 > 4.
[4] with maximum absolute diff |4-4| = 0 <= 4.
[4,7] with maximum absolute diff |4-7| = 3 <= 4.
[7] with maximum absolute diff |7-7| = 0 <= 4. 
Therefore, the size of the longest subarray is 2.
```

Example 2:

```
Input: nums = [10,1,2,4,7,2], limit = 5
Output: 4 
Explanation: The subarray [2,4,7,2] is the longest since the maximum absolute diff is |2-7| = 5 <= 5.
```

Example 3:

```
Input: nums = [4,2,2,2,4,4,2,2], limit = 0
Output: 3
```
 

Constraints:

1 <= nums.length <= 10^5  
1 <= nums[i] <= 10^9  
0 <= limit <= 10^9  


## 方法

```go
func longestSubarray(nums []int, limit int) int {
    minD := []int{} // queue for mins -> increasing order of mins from start to end
    maxD := []int{} // queue for maxs -> decresing order of maxs from start to end
    start := 0
    end := 0
    for end < len(nums) {
        current := nums[end]
        for len(minD) > 0 && last(minD) > current {
            minD = trim(minD)
        }
        for len(maxD) > 0 && last(maxD) < current {
            maxD = trim(maxD)
        }
        minD = append(minD, current)
        maxD = append(maxD, current)
        
        // Note: At this point, maxD has max, and minD has min numbers for start to end
        if len(maxD) > 0 && len(minD) > 0 && (maxD[0]-minD[0] > limit) {
            if maxD[0]==nums[start] {
                maxD=maxD[1:]
            }
            if minD[0]==nums[start] {
                minD=minD[1:]
            }
            start++
        }
        end++
    }
    return (end-start)
}

func last(a []int) int {
    return a[len(a)-1]
}

func trim(a []int) []int {
    return a[:len(a)-1]
}
```



```python
class Solution:
    def longestSubarray(self, nums: List[int], limit: int) -> int:
        minimal, maximal = float("inf"), float("-inf")
        size, current_size_start_number = 0, None
        for i in range(len(nums)):
            maximal = max(maximal, nums[i])
            minimal = min(minimal, nums[i])
            if abs(maximal - minimal) <= limit: 
                size += 1
            else:
                current_size_start_number = nums[i-size]
                if current_size_start_number == minimal:
                    minimal = min(nums[i - size + 1:i + 1])
                elif current_size_start_number == maximal:
                    maximal = max(nums[i - size + 1:i + 1])
        return size 
```