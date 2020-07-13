1512. Number of Good Pairs


Easy



Given an array of integers nums.

A pair (i,j) is called good if nums[i] == nums[j] and i < j.

Return the number of good pairs.

 

Example 1:

```
Input: nums = [1,2,3,1,1,3]
Output: 4
Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.
```

Example 2:

```
Input: nums = [1,1,1,1]
Output: 6
Explanation: Each pair in the array are good.
```

Example 3:

```
Input: nums = [1,2,3]
Output: 0
```
 

Constraints:

1 <= nums.length <= 100  
1 <= nums[i] <= 100

## 方法


```go
func numIdenticalPairs(nums []int) int {
    var index1, index2 int
	var r int
	for index1 < len(nums) && index2 < len(nums) {
		if index1 < index2 {
			if nums[index1] == nums[index2] {
				r++
			}
		}
		index2++
		if index2 == len(nums) {
			index1++
			index2 = 0
		}
		if index1 == len(nums)-1 {
			break
		}
	}
	return r
}
```



```python
class Solution:
    def numIdenticalPairs(self, A: List[int]) -> int:
        return sum(k * (k - 1) / 2 for k in collections.Counter(A).values())
```