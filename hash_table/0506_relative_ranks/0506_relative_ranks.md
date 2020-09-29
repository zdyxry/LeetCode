506. Relative Ranks


Easy


Given scores of N athletes, find their relative ranks and the people with the top three highest scores, who will be awarded medals: "Gold Medal", "Silver Medal" and "Bronze Medal".

Example 1:
```
Input: [5, 4, 3, 2, 1]
Output: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
Explanation: The first three athletes got the top three highest scores, so they got "Gold Medal", "Silver Medal" and "Bronze Medal". 
For the left two athletes, you just need to output their relative ranks according to their scores.
```

Note:  
N is a positive integer and won't exceed 10,000.  
All the scores of athletes are guaranteed to be unique.


## 方法



```go
func findRelativeRanks(nums []int) []string {
    list := make([]int, len(nums))
    copy(list, nums)
    sort.Slice(list, func (i int, j int) bool {
        return list[i] > list[j]
    })
    dict := make(map[int]int)
    for i, v := range list {
        dict[v] = i + 1
    }
    res := make([]string, len(nums))
    for i := 0; i < len(res); i++ {
        if dict[nums[i]] == 1 {
            res[i] = "Gold Medal"
        } else if dict[nums[i]] == 2 {
            res[i] = "Silver Medal"
        } else if dict[nums[i]] == 3 {
            res[i] = "Bronze Medal"
        } else {
            res[i] = strconv.Itoa(dict[nums[i]])
        }
    }
    return res
}
```

```python
class Solution:
    def findRelativeRanks(self, nums: List[int]) -> List[str]:
        hashmap = {}
        ans = [0] * len(nums)
        for i in range(len(nums)):
            hashmap[nums[i]] = i
        nums.sort(reverse = True)
        for i in range(len(nums)):
            if i == 0:
                ans[hashmap[nums[i]]] = 'Gold Medal'
            if i == 1:
                ans[hashmap[nums[i]]] = 'Silver Medal'
            if i == 2:
                ans[hashmap[nums[i]]] = 'Bronze Medal'
            if i > 2:
                ans[hashmap[nums[i]]] = str(i + 1)
        return ans
```