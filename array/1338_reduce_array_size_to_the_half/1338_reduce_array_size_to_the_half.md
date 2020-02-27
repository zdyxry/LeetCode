1338. Reduce Array Size to The Half


Medium


Given an array arr.  You can choose a set of integers and remove all the occurrences of these integers in the array.

Return the minimum size of the set so that at least half of the integers of the array are removed.

 

Example 1:

```
Input: arr = [3,3,3,3,5,5,5,2,2,7]
Output: 2
Explanation: Choosing {3,7} will make the new array [5,5,5,2,2] which has size 5 (i.e equal to half of the size of the old array).
Possible sets of size 2 are {3,5},{3,2},{5,2}.
Choosing set {2,7} is not possible as it will make the new array [3,3,3,3,5,5,5] which has size greater than half of the size of the old array.
```

Example 2:

```
Input: arr = [7,7,7,7,7,7]
Output: 1
Explanation: The only possible set you can choose is {7}. This will make the new array empty.
```

Example 3:

```
Input: arr = [1,9]
Output: 1
```

Example 4:

```
Input: arr = [1000,1000,3,7]
Output: 1
```

Example 5:

```
Input: arr = [1,2,3,4,5,6,7,8,9,10]
Output: 5
```
 

Constraints:

1 <= arr.length <= 10^5  
arr.length is even.  
1 <= arr[i] <= 10^5  

## 方法

```go
func minSetSize(arr []int) int {
    m := make(map[int]int)
    var r []int
    
    for _, v := range arr {
        m[v]++
    }
    
    for _, v := range m {
        r = append(r, v)
    }
    
    sort.Sort(sort.Reverse(sort.IntSlice(r)))
    
    var count int
    for idx, v := range r {
        count += v
        
        if count >= len(arr)/2 {
           return idx+1 
        }
    }
    
    return -1
}
```




```python
class Solution(object):
    def minSetSize(self, A):
        """
        :type arr: List[int]
        :rtype: int
        """
        import collections
        res, temp, n = 0, 0, len(A) / 2
        for a, freq in collections.Counter(A).most_common():
            temp += freq  # the number of the elements which should be removed
            res += 1
            if temp >= n:  # at least half of the integers of the array are removed
                return res
        return res
```