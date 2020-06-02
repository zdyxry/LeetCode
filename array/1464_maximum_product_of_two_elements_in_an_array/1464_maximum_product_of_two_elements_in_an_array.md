1464. Maximum Product of Two Elements in an Array


Easy


Given the array of integers nums, you will choose two different indices i and j of that array. Return the maximum value of (nums[i]-1)*(nums[j]-1).
 

Example 1:

```
Input: nums = [3,4,5,2]
Output: 12 
Explanation: If you choose the indices i=1 and j=2 (indexed from 0), you will get the maximum value, that is, (nums[1]-1)*(nums[2]-1) = (4-1)*(5-1) = 3*4 = 12. 
```

Example 2:

```
Input: nums = [1,5,4,5]
Output: 16
Explanation: Choosing the indices i=1 and j=3 (indexed from 0), you will get the maximum value of (5-1)*(5-1) = 16.
```

Example 3:

```
Input: nums = [3,7]
Output: 12
```
 

Constraints:

2 <= nums.length <= 500  
1 <= nums[i] <= 10^3


## 方法

```go
func maxProduct(nums []int) int {
    max:=0
    secondMax:=0
    for i:=0;i<len(nums);i++{
        e:=nums[i]
        if e>max{
            secondMax = max
            max = e
        }else if e>secondMax{
            secondMax = e
        }
    }
    return (max-1)*(secondMax-1)
}
```



```python
class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        nums=sorted(nums)
        return ((nums[-2]-1)*(nums[-1]-1))
```