462. Minimum Moves to Equal Array Elements II


Medium


Given a non-empty integer array, find the minimum number of moves required to make all array elements equal, where a move is incrementing a selected element by 1 or decrementing a selected element by 1.

You may assume the array's length is at most 10,000.

Example:

```
Input:
[1,2,3]

Output:
2

Explanation:
Only two moves are needed (remember each move increments or decrements one element):

[1,2,3]  =>  [2,2,3]  =>  [2,2,2]
```


## 方法




```go
func minMoves2(nums []int) int {
    lenth := len(nums)
    sort.Ints(nums)
    num := nums[lenth/2]
    res := 0
    for i:=0;i<lenth;i++{
        res+=getAbs(num,nums[i])
    }
    return res
}
func getAbs(a,b int)int{
    if a > b{
        return a - b
    }
    return b-a
}

```


```python
class Solution:
    def minMoves2(self, nums: List[int]) -> int:
        nums.sort()
        i, j = 0, len(nums)-1
        ans = 0
        while i<j:
            ans += nums[j]-nums[i]
            i += 1
            j -= 1     
        return ans
```