442. Find All Duplicates in an Array


Medium


Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements that appear twice in this array.

Could you do it without extra space and in O(n) runtime?

Example:
```
Input:
[4,3,2,7,8,2,3,1]

Output:
[2,3]
```


## 方法


```go
func findDuplicates(a []int) []int {
    for i := 0; i < len(a); i++ {
		for a[i] != a[a[i]-1] {
			a[i], a[a[i]-1] = a[a[i]-1], a[i]
		}
	}

	res := make([]int, 0, len(a))

	for i, n := range a {
		if i != n-1 {
			res = append(res, n)
		}
	}

	sort.Ints(res)

	return res
}
```


```python
class Solution(object):
    def findDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        result = []
        for i in nums:
            if nums[abs(i)-1] < 0:
                result.append(abs(i))
            else:
                nums[abs(i)-1] *= -1
        return result
```