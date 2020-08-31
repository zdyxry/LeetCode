1567. Maximum Length of Subarray With Positive Product


Medium


Given an array of integers nums, find the maximum length of a subarray where the product of all its elements is positive.

A subarray of an array is a consecutive sequence of zero or more values taken out of that array.

Return the maximum length of a subarray with positive product.

 

Example 1:

```
Input: nums = [1,-2,-3,4]
Output: 4
Explanation: The array nums already has a positive product of 24.
```

Example 2:

```
Input: nums = [0,1,-2,-3,-4]
Output: 3
Explanation: The longest subarray with positive product is [1,-2,-3] which has a product of 6.
Notice that we cannot include 0 in the subarray since that'll make the product 0 which is not positive.
```

Example 3:

```
Input: nums = [-1,-2,-3,0,1]
Output: 2
Explanation: The longest subarray with positive product is [-1,-2] or [-2,-3].
```

Example 4:

```
Input: nums = [-1,2]
Output: 1
```

Example 5:

```
Input: nums = [1,2,3,5,-6,4,0,10]
Output: 4
```
 

Constraints:

1 <= nums.length <= 10^5  
-10^9 <= nums[i] <= 10^9


## 方法

```go
func max(a, b int) int {
    if a > b { return a }
    return b
}
func getMaxLen(nums []int) int {
    odd, even, cnt, res := make([]int, 0), []int{-1}, 0, 0
    for i := range nums {
        if nums[i] == 0 {
            even = []int{i}
            odd = []int{}
            cnt = 0
        }
        if nums[i] < 0 {
            cnt++
            if len(odd) == 0 { odd = append(odd, i) }
        }
        if cnt%2 == 1 {
            res = max(res, i-odd[0])
        } else {
            res = max(res, i-even[0])
        }
    }
    return res
}
```



```python
class Solution:
    def getMaxLen(self, nums: List[int]) -> int:
        pre=-1
        l=[]
        res=0
        for i,num in enumerate(nums):
            if num<0:
                l.append(i)
            elif num==0:
                l,pre = [],i
            if len(l)%2==0:
                res=max(res,i-pre)
            else:
                res = max(res,i-l[0])
        return res
```