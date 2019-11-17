767. Reorganize String


Medium


Given a string S, check if the letters can be rearranged so that two characters that are adjacent to each other are not the same.

If possible, output any possible result.  If not possible, return the empty string.

Example 1:

```
Input: S = "aab"
Output: "aba"
```

Example 2:

```
Input: S = "aaab"
Output: ""
```

Note:

S will consist of lowercase letters and have length in range [1, 500].


## 方法


```go
func reorganizeString(s string) string {
    n := len(s)
	bs := []byte(s)

	count := make([][2]int, 26)
	maxCount := 0
	for _, b := range bs {
		count[b-'a'][0]++              // 0 位放字母的个数
		count[b-'a'][1] = int(b - 'a') // 1 位字母本身
		maxCount = max(maxCount, count[b-'a'][0])
	}

	// 利用个数最多的元素，进行快速判断。
	if 2*maxCount > n+1 {
		return ""
	}

	sort.Slice(count, func(i, j int) bool {
		if count[i][0] == count[j][0] {
			// 字母个数相同的时候，ascii 码小的在前面
			return count[i][1] < count[j][1]
		}
		// 个数多的字母排在前面
		return count[i][0] > count[j][0]
	})

	idx := 0

	for i := 0; i < 26 && count[i][0] > 0; i++ {
		b := byte('a' + count[i][1])

		for count[i][0] > 0 {
			bs[idx] = b
			count[i][0]--
			// 填写的时候，留下间隔
			// 避免相同的字母相邻
			idx += 2
			if idx >= n {
				// 到头了以后，从 1 位重新填写
				idx = 1
			}
		}

	}

	return string(bs)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```



```python
import collections
import heapq


class Solution(object):
    def reorganizeString(self, S):
        """
        :type S: str
        :rtype: str
        """
        counts = collections.Counter(S)
        if any(v > (len(S)+1)/2 for k, v in counts.iteritems()):
            return ""

        result = []
        max_heap = []
        for k, v in counts.iteritems():
            heapq.heappush(max_heap, (-v, k))
        while len(max_heap) > 1:
            count1, c1 = heapq.heappop(max_heap)
            count2, c2 = heapq.heappop(max_heap)
            if not result or c1 != result[-1]:
                result.extend([c1, c2])
                if count1+1: heapq.heappush(max_heap, (count1+1, c1))
                if count2+1: heapq.heappush(max_heap, (count2+1, c2))
        return "".join(result) + (max_heap[0][1] if max_heap else '')
```