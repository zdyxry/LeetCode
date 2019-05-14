88. Merge Sorted Array

Easy

Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
Example:

```
Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]
```


# 方法
从后向前遍历，取大数放置在 nums1 的最后，依次向前移动。当 nums1 还存在数字未合并时，无须操作，nums1 和 nums2 都是已经排序的数组，故前置数字肯定比后置数字小，保持原有位置即可；单独对 nums2 处理。




```go
func merge(nums1 []int, m int, nums2 []int, n int)  {
    i := m - 1;
	j := n - 1;
    for i >= 0 && j >= 0 {
		if nums1[i] <= nums2[j] {
			nums1[i + j + 1] = nums2[j]
			j --
		} else {
			nums1[i + j + 1] = nums1[i]
			i --
		}
        
	}
    for j >= 0 {
        nums1[i + j + 1] = nums2[j]
        j --
    }
    fmt.Println(i, j)
}
```



```python
class Solution(object):
    # @param nums1  a list of integers
    # @param m  an integer, length of nums1
    # @param nums2  a list of integers
    # @param n  an integer, length of nums2
    # @return nothing
    def merge(self, nums1, m, nums2, n):
        last, i, j = m + n - 1, m - 1, n - 1

        while i >= 0 and j >= 0:
            if nums1[i] > nums2[j]:
                nums1[last] = nums1[i]
                last, i = last - 1, i - 1
            else:
                nums1[last] = nums2[j]
                last, j = last - 1, j - 1
        print(i, j, last)
        while j >= 0:
            nums1[last] = nums2[j]
            last, j = last - 1, j - 1
```