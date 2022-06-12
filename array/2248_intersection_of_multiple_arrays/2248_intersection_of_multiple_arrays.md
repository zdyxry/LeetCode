2248. Intersection of Multiple Arrays


Easy

Given a 2D integer array nums where nums[i] is a non-empty array of distinct positive integers, return the list of integers that are present in each array of nums sorted in ascending order.
 

Example 1:

```
Input: nums = [[3,1,2,4,5],[1,2,3,4],[3,4,5,6]]
Output: [3,4]
Explanation: 
The only integers present in each of nums[0] = [3,1,2,4,5], nums[1] = [1,2,3,4], and nums[2] = [3,4,5,6] are 3 and 4, so we return [3,4].
```

Example 2:

```
Input: nums = [[1,2,3],[4,5,6]]
Output: []
Explanation: 
There does not exist any integer present both in nums[0] and nums[1], so we return an empty list [].
 ``

Constraints:

```
1 <= nums.length <= 1000
1 <= sum(nums[i].length) <= 1000
1 <= nums[i][j] <= 1000
All the values of nums[i] are unique.
```


## 方法



```
func intersection(nums [][]int) (ans []int) {
	cnt := map[int]int{}
	for _, a := range nums {
		for _, v := range a {
			cnt[v]++
		}
	}
	for v, c := range cnt {
		if c == len(nums) {
			ans = append(ans, v)
		}
	}
	sort.Ints(ans)
	return
}

```


```python
class Solution:
    def intersection(self, nums: List[List[int]]) -> List[int]:
        result = nums[0]
        for arr in nums[1:]:
            result = [i for i in result if i in arr]
        result.sort()
        return result
```