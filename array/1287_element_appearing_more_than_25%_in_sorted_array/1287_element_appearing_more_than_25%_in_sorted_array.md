1287. Element Appearing More Than 25% In Sorted Array


Easy


Given an integer array sorted in non-decreasing order, there is exactly one integer in the array that occurs more than 25% of the time.

Return that integer.

 

Example 1:

```
Input: arr = [1,2,2,6,6,6,6,7,10]
Output: 6
```
 

Constraints:

1 <= arr.length <= 10^4  
0 <= arr[i] <= 10^5

## 方法

```go
func findSpecialInteger(arr []int) int {
    length := len(arr)
    cur := arr[0]
    cnt := 0
    for _, v := range arr {
        if v == cur {
            cnt++
            if cnt * 4 > length {
                return v
            }
        } else {
            cur = v
            cnt = 1
        }
    }
    return -1
    
}
```



```python
class Solution(object):
    def findSpecialInteger(self, arr):
        """
        :type arr: List[int]
        :rtype: int
        """
        n = len(arr)
        cur, cnt = arr[0], 0
        for i in range(n):
            if arr[i] == cur:
                cnt += 1
                if cnt * 4 > n:
                    return cur
            else:
                cur, cnt = arr[i], 1
        return -1

```