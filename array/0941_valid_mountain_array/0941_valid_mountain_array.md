941. Valid Mountain Array


Easy


Given an array A of integers, return true if and only if it is a valid mountain array.

Recall that A is a mountain array if and only if:

```
A.length >= 3
There exists some i with 0 < i < A.length - 1 such that:
A[0] < A[1] < ... A[i-1] < A[i]
A[i] > A[i+1] > ... > A[A.length - 1]
```

 

Example 1:

```
Input: [2,1]
Output: false
```

Example 2:

```
Input: [3,5,5]
Output: false
```

Example 3:

```
Input: [0,3,2,1]
Output: true
```

Note:

0 <= A.length <= 10000  
0 <= A[i] <= 10000 


## 方法


```go
func validMountainArray(A []int) bool {
    length:=len(A)
    if length < 3 {
        return false
    }
    count:=0 // 转折点 本题山脉应当是 左上坡 右下坡 且 只有一个转折点
    pre:=0 // 前一个点的趋势 1为上坡 -1为下坡
    for i:=0; i<length-1; i++ {
        if A[i] == A[i+1] {
            return false
        }
        if A[i] > A[i+1] && i == 0 {
            return false
        }
        if A[i] < A[i+1] { // 上坡
            if pre == -1 {
                return false
            }
            pre=1
        }
        if A[i] > A[i+1] { // 下坡
            if pre == 1 {
                count ++
            }
            pre = -1
        }
    }
    return count == 1
}
```


```python
class Solution:
    def validMountainArray(self, A: List[int]) -> bool:
        l,r=0,len(A)-1
        while l<r and A[l]<A[l+1]: l+=1
        while r>l and A[r]<A[r-1]: r-=1
        return l==r and l!=0 and r!=len(A)-1
```