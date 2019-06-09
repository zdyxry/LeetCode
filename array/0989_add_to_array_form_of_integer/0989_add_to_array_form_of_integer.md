989. Add to Array-Form of Integer

Easy

For a non-negative integer X, the array-form of X is an array of its digits in left to right order.  For example, if X = 1231, then the array form is [1,2,3,1].

Given the array-form A of a non-negative integer X, return the array-form of the integer X+K.

 

Example 1:

```
Input: A = [1,2,0,0], K = 34
Output: [1,2,3,4]
Explanation: 1200 + 34 = 1234
```
Example 2:

```
Input: A = [2,7,4], K = 181
Output: [4,5,5]
Explanation: 274 + 181 = 455
```
Example 3:

```
Input: A = [2,1,5], K = 806
Output: [1,0,2,1]
Explanation: 215 + 806 = 1021
```
Example 4:

```
Input: A = [9,9,9,9,9,9,9,9,9,9], K = 1
Output: [1,0,0,0,0,0,0,0,0,0,0]
Explanation: 9999999999 + 1 = 10000000000
```
 

Note：

1 <= A.length <= 10000

0 <= A[i] <= 9

0 <= K <= 10000

If A.length > 1, then A[0] != 0


# 方法
对 A 最后一位与 K 做加法，判断是否需要进位。


```go

func addToArrayForm(A []int, K int) []int {
	size := len(A)
	A[size-1] += K
	for i := size - 1; i > 0 && A[i] > 9; i-- {
		A[i-1] += A[i] / 10
		A[i] %= 10
	}

	if A[0] < 10 {
		return A
	}

	A0 := num2ints(A[0])
	return append(A0, A[1:]...)
}

func num2ints(n int) []int {
	res := make([]int, 0, 8)
	for n > 0 {
		res = append(res, n%10)
		n /= 10
	}
	reverse(res)
	return res
}

func reverse(A []int) {
	i, j := 0, len(A)-1
	for i < j {
		A[i], A[j] = A[j], A[i]
		i++
		j--
	}
}


```

```python
class Solution(object):
    def addToArrayForm(self, A, K):
        """
        :type A: List[int]
        :type K: int
        :rtype: List[int]
        """
        A.reverse()
        carry, i = K, 0
        A[i] += carry
        carry, A[i] = divmod(A[i], 10)
        while carry:
            i += 1
            if i < len(A):
                A[i] += carry
            else:
                A.append(carry)
            carry, A[i] = divmod(A[i], 10)
        A.reverse()
        return A

```