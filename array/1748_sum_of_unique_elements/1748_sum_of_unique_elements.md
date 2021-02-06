1748. Sum of Unique Elements


Easy


You are given an integer array nums. The unique elements of an array are the elements that appear exactly once in the array.

Return the sum of all the unique elements of nums.

 

Example 1:

```
Input: nums = [1,2,3,2]
Output: 4
Explanation: The unique elements are [1,3], and the sum is 4.
```

Example 2:

```
Input: nums = [1,1,1,1,1]
Output: 0
Explanation: There are no unique elements, and the sum is 0.
```

Example 3:

```
Input: nums = [1,2,3,4,5]
Output: 15
Explanation: The unique elements are [1,2,3,4,5], and the sum is 15.
```

Constraints:

1 <= nums.length <= 100   
1 <= nums[i] <= 100


## 方法


```go
func sumOfUnique(nums []int) int {
    counter := make(map[int]int)
    for _, x := range nums {
        if _, ok := counter[x]; !ok {
            counter[x] = 1
        } else {
            counter[x]++
        }
    }
    res := 0
    for k, v := range counter {
        if v == 1 {
            res += k
        }
    }
    return res
}
```


```python
class Solution:
    def sumOfUnique(self, nums: List[int]) -> int:
        import collections
        return sum([i for i,v in collections.Counter(nums).items() if v <= 1])
```