453. Minimum Moves to Equal Array Elements


Easy


Given a non-empty integer array of size n, find the minimum number of moves required to make all array elements equal, where a move is incrementing n - 1 elements by 1.

Example:

```
Input:
[1,2,3]

Output:
3

Explanation:
Only three moves are needed (remember each move increments two elements):

[1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
```


## 方法


```go
func minMoves(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    minSoFar :=  nums[0]
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if minSoFar > nums[i] {
            minSoFar = nums[i]
        }
    }
    return sum - len(nums)*minSoFar
}
```


```python
class Solution:
    def minMoves(self, nums: List[int]) -> int:
        sum = 0
        minmum = min(nums)
        for i in nums:
            sum += i-minmum
        return sum
```