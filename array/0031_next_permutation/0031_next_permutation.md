31. Next Permutation


Medium


Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

```
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
```

## 方法

```go
func nextPermutation(a []int)  {
    left := len(a) - 2
	for 0 <= left && a[left] >= a[left+1] {
		left--
	}

	// 此时 a[left+1:] 是一个 递减 数列

	reverse(a, left+1)

	if left == -1 {
		return
	}

	// 此时 a[left+1:] 是一个 递增 数列

	right := search(a, left+1, a[left])
	a[left], a[right] = a[right], a[left]
}

// 逆转 a[l:]
func reverse(a []int, l int) {
	r := len(a) - 1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}

// 返回 a[l:] 中 > target 的最小值的索引号
// a[l:] 是一个 递增 数列
func search(a []int, l, target int) int {
	r := len(a) - 1
	l--
	for l+1 < r {
		mid := (l + r) / 2
		if target < a[mid] {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

```


```python
class Solution(object):
    def nextPermutation(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        i = j = len(nums)-1
        while i > 0 and nums[i-1] >= nums[i]:
            i -= 1
        if i == 0:   # nums are in descending order
            nums.reverse()
            return 
        k = i - 1    # find the last "ascending" position
        while nums[j] <= nums[k]:
            j -= 1
        nums[k], nums[j] = nums[j], nums[k]  
        l, r = k+1, len(nums)-1  # reverse the second part
        while l < r:
            nums[l], nums[r] = nums[r], nums[l]
            l +=1 ; r -= 1
```