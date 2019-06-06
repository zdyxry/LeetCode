905. Sort Array By Parity

Easy

Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.

You may return any answer array that satisfies this condition.


Example 1:

```
Input: [3,1,2,4]
Output: [2,4,3,1]
The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
```
 

Note:

1 <= A.length <= 5000

0 <= A[i] <= 5000


# 方法
双指针，前者指向偶数位置，后者指向奇数位置， 依次遍历数组，分别添加到对应位置中。


```go
func sortArrayByParity(A []int) []int {
    left := 0
	right := len(A) - 1

	res := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			res[left] = A[i]
			left++
		} else {
			res[right] = A[i]
			right--
		}
	}

	return res
    
}
```


```python
class Solution(object):
    def sortArrayByParity(self, A):
        """
        :type A: List[int]
        :rtype: List[int]
        """
        if not A:
            return []
        
        start, end = 0, len(A)-1
        
        while (start<end):
            while (start<end and A[start]%2==0):
                start+=1
            while (start<end and A[end]%2==1):
                end-=1
            
            if start<end:
                A[start],A[end]=A[end],A[start]
                start+=1
                end-=1
        
        return A
```