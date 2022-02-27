2148. Count Elements With Strictly Smaller and Greater Elements


Easy


Given an integer array nums, return the number of elements that have both a strictly smaller and a strictly greater element appear in nums.

 

Example 1:

```
Input: nums = [11,7,2,15]
Output: 2
Explanation: The element 7 has the element 2 strictly smaller than it and the element 11 strictly greater than it.
Element 11 has element 7 strictly smaller than it and element 15 strictly greater than it.
In total there are 2 elements having both a strictly smaller and a strictly greater element appear in nums.
```

Example 2:

```
Input: nums = [-3,3,3,90]
Output: 2
Explanation: The element 3 has the element -3 strictly smaller than it and the element 90 strictly greater than it.
Since there are two elements with the value 3, in total there are 2 elements having both a strictly smaller and a strictly greater element appear in nums.
```

Constraints:

1 <= nums.length <= 100   
-105 <= nums[i] <= 105


## 方法


```
func countElements(nums []int) int {
    sort.Ints(nums)
    
    result := 0
    for idx := range nums {
        if nums[0] < nums[idx]  && nums[idx] < nums[len(nums)-1] {
            result += 1
        }
    }
    return result
    
}
```


```
class Solution:
    def countElements(self, nums: List[int]) -> int:
        smallest = min(nums)
        largest = max(nums)
        res = 0
        for num in nums:
            if smallest < num < largest:
                res += 1
        return res

```