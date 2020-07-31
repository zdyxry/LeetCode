1524. Number of Sub-arrays With Odd Sum


Medium


Given an array of integers arr. Return the number of sub-arrays with odd sum.

As the answer may grow large, the answer must be computed modulo 10^9 + 7.

 

Example 1:

```
Input: arr = [1,3,5]
Output: 4
Explanation: All sub-arrays are [[1],[1,3],[1,3,5],[3],[3,5],[5]]
All sub-arrays sum are [1,4,9,3,8,5].
Odd sums are [1,9,3,5] so the answer is 4.
```

Example 2:

```
Input: arr = [2,4,6]
Output: 0
Explanation: All sub-arrays are [[2],[2,4],[2,4,6],[4],[4,6],[6]]
All sub-arrays sum are [2,6,12,4,10,6].
All sub-arrays have even sum and the answer is 0.
```

Example 3:

```
Input: arr = [1,2,3,4,5,6,7]
Output: 16
```

Example 4:

```
Input: arr = [100,100,99,99]
Output: 4
```

Example 5:

```
Input: arr = [7]
Output: 1
```
 

Constraints:

1 <= arr.length <= 10^5  
1 <= arr[i] <= 100

## 方法


```go

const mod = int(1e9 + 7)

func numOfSubarrays(arr []int) int {
    odd, even := 0, 0
    res := 0
    for _, v := range arr {
        if v % 2 != 0 {
            odd, even = even + 1, odd
        } else {
            odd, even = odd, even + 1
        }
        res += odd
        res %= mod
    }
    return res
}
```


```python
class Solution:
    def numOfSubarrays(self, arr: List[int]) -> int:
        pre = 0
        odd, even = 0, 1
        for num in arr:
            pre += num
            if pre%2 == 0:
                even +=1
            else:
                odd +=1
        return (even*odd)%(10**9 + 7)
```