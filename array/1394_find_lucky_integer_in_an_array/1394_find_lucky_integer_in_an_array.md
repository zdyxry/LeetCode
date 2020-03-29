1394. Find Lucky Integer in an Array


Easy


Given an array of integers arr, a lucky integer is an integer which has a frequency in the array equal to its value.

Return a lucky integer in the array. If there are multiple lucky integers return the largest of them. If there is no lucky integer return -1.

 

Example 1:

```
Input: arr = [2,2,3,4]
Output: 2
Explanation: The only lucky number in the array is 2 because frequency[2] == 2.
```

Example 2:

```
Input: arr = [1,2,2,3,3,3]
Output: 3
Explanation: 1, 2 and 3 are all lucky numbers, return the largest of them.
```

Example 3:

```
Input: arr = [2,2,2,3,3]
Output: -1
Explanation: There are no lucky numbers in the array.
```

Example 4:

```
Input: arr = [5]
Output: -1
```

Example 5:

```
Input: arr = [7,7,7,7,7,7,7]
Output: 7
```
 

Constraints:

1 <= arr.length <= 500  
1 <= arr[i] <= 500

## 方法

```go
func findLucky(arr []int) int {
    freqCounter := make(map[int]int)
    
    for _, elem := range arr {
        freqCounter[elem] += 1
    }
    result := -1
    for key, val := range freqCounter {
        if key == val && result < val {
            result = val
        }
    }
    return result
}
```



```python
class Solution(object):
    def findLucky(self, arr):
        """
        :type arr: List[int]
        :rtype: int
        """
        cnt = [0] * 501
        for a in arr:
            cnt[a] += 1
        for i in range(500, 0, -1):
            if cnt[i] == i:
                return i    
        return -1
```