2012. Sum of Beauty in the Array


Medium

You are given a 0-indexed integer array nums. For each index i (1 <= i <= nums.length - 2) the beauty of nums[i] equals:

```
2, if nums[j] < nums[i] < nums[k], for all 0 <= j < i and for all i < k <= nums.length - 1.
1, if nums[i - 1] < nums[i] < nums[i + 1], and the previous condition is not satisfied.
0, if none of the previous conditions holds.
Return the sum of beauty of all nums[i] where 1 <= i <= nums.length - 2.
```
 

Example 1:

```
Input: nums = [1,2,3]
Output: 2
Explanation: For each index i in the range 1 <= i <= 1:
- The beauty of nums[1] equals 2.
```

Example 2:

```
Input: nums = [2,4,6,4]
Output: 1
Explanation: For each index i in the range 1 <= i <= 2:
- The beauty of nums[1] equals 1.
- The beauty of nums[2] equals 0.
```

Example 3:

```
Input: nums = [3,2,1]
Output: 0
Explanation: For each index i in the range 1 <= i <= 1:
- The beauty of nums[1] equals 0.
```

Constraints:

3 <= nums.length <= 105   
1 <= nums[i] <= 105    


## 方法


```
func sumOfBeauties(nums []int) int {

	total := 0

	n := len(nums)

	pre := make([]int, n)
	suf := make([]int, n)

	pre[0] = nums[0]
	suf[n-1] = nums[n-1]

	for i := 1; i < n; i++ {
		if pre[i-1] < nums[i] {
			pre[i] = nums[i]
		} else {
			pre[i] = pre[i-1]
		}
	}

	for i := n - 2; i >= 0; i-- {
		if nums[i] < suf[i+1] {
			suf[i] = nums[i]
		} else {
			suf[i] = suf[i+1]
		}
	}

	fmt.Println(pre)
	fmt.Println(suf)

	for i := 1; i <= n-2; i++ {
		if pre[i-1] < nums[i] && nums[i] < suf[i+1] {
			total += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			total++
		}
	}

	return total
}

```


```class Solution:
    def sumOfBeauties(self, nums: List[int]) -> int:
        n = len(nums)
        ans = 0
        leftMax,rightMin = [0] * n, [0] * n
        leftMax[0] = nums[0]
        rightMin[-1] = nums[-1]
        # 计算左边的最大值
        for i in range(1, n):
            if nums[i] > leftMax[i-1]:
                leftMax[i] = nums[i]
            else:
                leftMax[i] = leftMax[i-1]
        # 计算右边最小值
        for j in range(n-2, -1, -1):
            if nums[j] < rightMin[j + 1]:
                rightMin[j] = nums[j]
            else:
                rightMin[j] = rightMin[j+1]
        # 判断美丽值
        for i in range(1,n-1):
            if leftMax[i-1] < nums[i] < rightMin[i+1]:
                ans += 2
            elif nums[i-1] < nums[i] < nums[i+1]:
                ans += 1
            else:
                pass
        return ans
```