1588. Sum of All Odd Length Subarrays


Easy


Given an array of positive integers arr, calculate the sum of all possible odd-length subarrays.

A subarray is a contiguous subsequence of the array.

Return the sum of all odd-length subarrays of arr.

 

Example 1:

```
Input: arr = [1,4,2,5,3]
Output: 58
Explanation: The odd-length subarrays of arr and their sums are:
[1] = 1
[4] = 4
[2] = 2
[5] = 5
[3] = 3
[1,4,2] = 7
[4,2,5] = 11
[2,5,3] = 10
[1,4,2,5,3] = 15
If we add all these together we get 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58
```

Example 2:

```
Input: arr = [1,2]
Output: 3
Explanation: There are only 2 subarrays of odd length, [1] and [2]. Their sum is 3.
```

Example 3:

```
Input: arr = [10,11,12]
Output: 66
```
 

Constraints:

1 <= arr.length <= 100   
1 <= arr[i] <= 1000


## 方法


```go
func sumOddLengthSubarrays(arr []int) int {
    if arr == nil || len(arr) == 0 {
        return 0
    }
    var count = 0
    var tmp =0 
    for i:= 0 ; i < len(arr) ; i ++ {
        tmp = arr[i]
        count = count + tmp
        for j := i + 1 ; j < len(arr) ; j ++ {
            tmp = tmp + arr[j]
            if (j - i) % 2 == 0 {
                count = count + tmp
            }
        }
    }
    return count
}
```


```python
class Solution:
    def sumOddLengthSubarrays(self, arr: List[int]) -> int:
        res = 0
        for i in range(len(arr)+1):
            if i % 2:
                if i == 1:
                    res += sum(arr)
                else:
                    j = 0
                    while j + i <= len(arr):
                        res += sum(arr[j:j+i])
                        j += 1
        return res
```