1343. Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold


Medium


Given an array of integers arr and two integers k and threshold.

Return the number of sub-arrays of size k and average greater than or equal to threshold.

 

Example 1:

```
Input: arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4
Output: 3
Explanation: Sub-arrays [2,5,5],[5,5,5] and [5,5,8] have averages 4, 5 and 6 respectively. All other sub-arrays of size 3 have averages less than 4 (the threshold).
```

Example 2:

```
Input: arr = [1,1,1,1,1], k = 1, threshold = 0
Output: 5
```

Example 3:

```
Input: arr = [11,13,17,23,29,31,7,5,2,3], k = 3, threshold = 5
Output: 6
Explanation: The first 6 sub-arrays of size 3 have averages greater than 5. Note that averages are not integers.
```


Example 4:

```
Input: arr = [7,7,7,7,7,7,7], k = 7, threshold = 7
Output: 1
```

Example 5:

```
Input: arr = [4,4,4,4], k = 4, threshold = 1
Output: 1
```

## 方法

```go
func numOfSubarrays(arr []int, k int, threshold int) int {
    var (
        i, sum, count int
        length        = len(arr)
    )

    for i = 0; i < k; i++ {
        sum += arr[i]
    }

    for {
        if threshold*k <= sum {
            count++
        }
        if i >= length {
            break
        }
        sum = sum - arr[i-k] + arr[i]
        i++
    }
    return count
}
```


```python
class Solution(object):
    def numOfSubarrays(self, arr, k, threshold):
        """
        :type arr: List[int]
        :type k: int
        :type threshold: int
        :rtype: int
        """
        n = len(arr)
        left = 0
        res = 0
        for i in range(n):
            left += arr[i]
            if i >= k-1:
                if left >= threshold * k:
                    res += 1
                left -= arr[i-k+1]
        return res
```
