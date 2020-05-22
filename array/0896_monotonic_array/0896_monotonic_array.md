896. Monotonic Array


Easy


An array is monotonic if it is either monotone increasing or monotone decreasing.

An array A is monotone increasing if for all i <= j, A[i] <= A[j].  An array A is monotone decreasing if for all i <= j, A[i] >= A[j].

Return true if and only if the given array A is monotonic.

 

Example 1:

```
Input: [1,2,2,3]
Output: true
```

Example 2:

```
Input: [6,5,4,4]
Output: true
```

Example 3:

```
Input: [1,3,2]
Output: false
```

Example 4:

```
Input: [1,2,4,5]
Output: true
```

Example 5:

```
Input: [1,1,1]
Output: true
```
 

Note:

1 <= A.length <= 50000  
-100000 <= A[i] <= 100000


## 方法

```go
func isMonotonic(A []int) bool {
    if len(A) <= 2 {return true}
    noInc, noDec := true, true
    for i:=0; i<len(A)-1; i++{
        if A[i] > A[i+1] {
            noDec = false
        }
        if A[i] < A[i+1] {
            noInc = false
        }
        
    }
    return noDec || noInc
}
```



```python
class Solution:
    def isMonotonic(self, A: List[int]) -> bool:
        cnt_inc = 0; cnt_dec = 0
        for i, v in enumerate(A):
            if i == 0: prev = v; continue
            if   v > prev: cnt_inc += 1
            elif v < prev: cnt_dec += 1
            if cnt_inc and cnt_dec: return False
            prev = v
        return True
```