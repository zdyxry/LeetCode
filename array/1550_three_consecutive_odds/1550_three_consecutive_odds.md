1550. Three Consecutive Odds


Easy


Given an integer array arr, return true if there are three consecutive odd numbers in the array. Otherwise, return false.
 

Example 1:

```
Input: arr = [2,6,4,1]
Output: false
Explanation: There are no three consecutive odds.
```

Example 2:

```
Input: arr = [1,2,34,3,4,5,7,23,12]
Output: true
Explanation: [5,7,23] are three consecutive odds.
```

Constraints:

1 <= arr.length <= 1000  
1 <= arr[i] <= 1000


## 方法

```go
func threeConsecutiveOdds(arr []int) bool {
    cnt := 0
    for _, n := range arr {
        if n & 1 == 1{
            cnt++
        } else {
            cnt = 0
        }
        if cnt == 3 {
            return true
        }
    }
    return false
    
}
```


```python
class Solution:
    def threeConsecutiveOdds(self, arr: List[int]) -> bool:
        cnt = 0
        for i in arr:
            if i & 1:
                cnt += 1
            else:
                cnt = 0
            if cnt == 3:
                return True
        return False
```