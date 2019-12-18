659. Split Array into Consecutive Subsequences


Medium


Given an array nums sorted in ascending order, return true if and only if you can split it into 1 or more subsequences such that each subsequence consists of consecutive integers and has length at least 3.

 

Example 1:

```
Input: [1,2,3,3,4,5]
Output: True
Explanation:
You can split them into two consecutive subsequences : 
1, 2, 3
3, 4, 5
```

Example 2:

```
Input: [1,2,3,3,4,4,5,5]
Output: True
Explanation:
You can split them into two consecutive subsequences : 
1, 2, 3, 4, 5
3, 4, 5
```


Example 3:

```
Input: [1,2,3,4,4,5]
Output: False
```
 

Constraints:

1 <= nums.length <= 10000


## 方法

```go
func isPossible(nums []int) bool {
    // 用于存放每个数字出现的次数
    counter := make(map[int]int)
    // 用于存放子序列结尾的那个数
    tail := make(map[int]int)
    for i := 0; i < len(nums); i++{
        counter[nums[i]]++
    }
    
    for _, v := range nums {
        // 先判断数字是否被消耗完
        if counter[v] == 0 {
            continue
        }
        // 判断是否能接着之前子序列
        if n := tail[v-1]; n > 0{
            tail[v-1]--
            tail[v]++
            counter[v]--
            continue
        }
        // 判断是否能自己组成新的子序列
        if counter[v+1] > 0 && counter[v+2] > 0 {
            tail[v+2]++
            counter[v]--
            counter[v+1]--
            counter[v+2]--
        } else {
            return false
        }
    }
    return true

}
```



```python
class Solution(object):
    def isPossible(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        pre, cur = float("-inf"), 0
        cnt1, cnt2, cnt3 = 0, 0, 0
        i = 0
        while i < len(nums):
            cnt = 0
            cur = nums[i]
            while i < len(nums) and cur == nums[i]:
                cnt += 1
                i += 1

            if cur != pre + 1:
                if cnt1 != 0 or cnt2 != 0:
                    return False
                cnt1, cnt2, cnt3 = cnt, 0, 0
            else:
                if cnt < cnt1 + cnt2:
                    return False
                cnt1, cnt2, cnt3 = max(0, cnt - (cnt1 + cnt2 + cnt3)), \
                                   cnt1, \
                                   cnt2 + min(cnt3, cnt - (cnt1 + cnt2))
            pre = cur
        return cnt1 == 0 and cnt2 == 0
```

```python
class Solution(object):
    def isPossible(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        saved = collections.defaultdict(list)
        for num in nums:
            last = saved[num - 1]
            _len = 0 if (not last) else heapq.heappop(last)
            current = saved[num]
            heapq.heappush(current, _len + 1)
        for values in saved.values():
            for v in values:
                if v < 3:
                    return False
        return True
```