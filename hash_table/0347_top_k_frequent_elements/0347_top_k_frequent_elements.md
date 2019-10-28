347. Top K Frequent Elements


Medium


Given a non-empty array of integers, return the k most frequent elements.

Example 1:

```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

Example 2:

```
Input: nums = [1], k = 1
Output: [1]
```

Note:

You may assume k is always valid, 1 ≤ k ≤ number of unique elements.

Your algorithm's time complexity must be better than O(n log n), where n is the array's size.

## 方法

```go
func topKFrequent(nums []int, k int) []int {
    res := make([]int, 0, k)
	// 统计每个数字出现的次数
	rec := make(map[int]int, len(nums))
	for _, n := range nums {
		rec[n]++
	}
	// 对出现次数进行排序
	counts := make([]int, 0, len(rec))
	for _, c := range rec {
		counts = append(counts, c)
	}
	sort.Ints(counts)
	// min 是 前 k 个高频数字的底线
	min := counts[len(counts)-k]

	// 收集所有　不低于　底线的数字
	for n, c := range rec {
		if c >= min {
			res = append(res, n)
		}
	}

	return res
}
```


```python
class Solution(object):
    def topKFrequent(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: List[int]
        """
        hs = {}
        frq = {}
        for i in xrange(0, len(nums)):
            if nums[i] not in hs:
                hs[nums[i]] = 1
            else:
                hs[nums[i]] += 1

        for z,v in hs.iteritems():
            if v not in frq:
                frq[v] = [z]
            else:
                frq[v].append(z)
        
        arr = []
        
        for x in xrange(len(nums), 0, -1):
            if x in frq:
                
                for i in frq[x]:
                    arr.append(i)

        return [arr[x] for x in xrange(0, k)]
```