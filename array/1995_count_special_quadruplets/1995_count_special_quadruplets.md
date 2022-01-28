1995. Count Special Quadruplets


Easy

Given a 0-indexed integer array nums, return the number of distinct quadruplets (a, b, c, d) such that:

nums[a] + nums[b] + nums[c] == nums[d], and a < b < c < d
 

Example 1:

```
Input: nums = [1,2,3,6]
Output: 1
Explanation: The only quadruplet that satisfies the requirement is (0, 1, 2, 3) because 1 + 2 + 3 == 6.
```

Example 2:

```
Input: nums = [3,3,6,4,5]
Output: 0
Explanation: There are no such quadruplets in [3,3,6,4,5].
```

Example 3:

```
Input: nums = [1,1,1,3,5]
Output: 4
Explanation: The 4 quadruplets that satisfy the requirement are:
- (0, 1, 2, 3): 1 + 1 + 1 == 3
- (0, 1, 3, 4): 1 + 1 + 3 == 5
- (0, 2, 3, 4): 1 + 1 + 3 == 5
- (1, 2, 3, 4): 1 + 1 + 3 == 5
```

Constraints:

4 <= nums.length <= 50   
1 <= nums[i] <= 100

## 方法



```
func countQuadruplets(nums []int) (ans int) {
    cnts := map[int]int{}
    for i := 1; i < len(nums) - 2; i++ {
        for j := 0; j < i; j++{
            cnts[nums[i] + nums[j]]++
        }
        for j := i + 2; j < len(nums); j++ {
            ans += cnts[nums[j] - nums[i+1]]
        }
    }
    return
}

```



```
class Solution:
    def countQuadruplets(self, nums: List[int]) -> int:
        if len(nums) < 4:
            return 0
            
        result = 0
        for i in range(0, len(nums)):
            for j in range(i+1, len(nums)):
                for m in range(j+1, len(nums)):
                    for n in range(m+1, len(nums)):
                        if nums[i] + nums[j] + nums[m] == nums[n]:
                            result += 1
        return result

```