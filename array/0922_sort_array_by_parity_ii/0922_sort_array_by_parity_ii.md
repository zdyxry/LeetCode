922. Sort Array By Parity II

Easy

Given an array A of non-negative integers, half of the integers in A are odd, and half of the integers are even.


Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is even, i is even.


You may return any answer array that satisfies this condition.

 

Example 1:

```
Input: [4,2,5,7]
Output: [4,5,2,7]
Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.
```
 

Note:

2 <= A.length <= 20000

A.length % 2 == 0

0 <= A[i] <= 1000


# 方法
双指针遍历数组，当当前数值为奇数时，跟后续下标为奇数但数字为偶数的数字做交换。
比如：

4,2,5,7

4 为偶数，略过；

5 为奇数， 则找到最后一个下标为奇数但数字为偶数的数字： A[1] = 2，进行交换；

4,5,2,7





```go
func sortArrayByParityII(A []int) []int {
    for i, j := 0, 1; i < len(A); i += 2 {
        if A[i] % 2 == 1 {
            for A[j] % 2 == 1 {
                j += 2
            }
            A[i], A[j] = A[j], A[i]
        }
    }
    return A
}
```


```python
class Solution(object):
    def sortArrayByParityII(self, A):
        """
        :type A: List[int]
        :rtype: List[int]
        """
        j = 1
        for i in xrange(0, len(A), 2):
            if A[i] % 2:
                while A[j] % 2:
                    j += 2
                A[i], A[j] = A[j], A[i]
            print(A)
        return A

```