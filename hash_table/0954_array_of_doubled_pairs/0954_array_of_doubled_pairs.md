954. Array of Doubled Pairs


Medium


Given an array of integers A with even length, return true if and only if it is possible to reorder it such that A[2 * i + 1] = 2 * A[2 * i] for every 0 <= i < len(A) / 2.

 

Example 1:

```
Input: [3,1,3,6]
Output: false
```

Example 2:

```
Input: [2,1,2,6]
Output: false
```

Example 3:

```
Input: [4,-2,2,-4]
Output: true
Explanation: We can take two groups, [-2,-4] and [2,4] to form [-2,-4,2,4] or [2,4,-2,-4].
```

Example 4:
```
Input: [1,2,4,16,8,4]
Output: false
```
 

Note:

0 <= A.length <= 30000  
A.length is even  
-100000 <= A[i] <= 100000  

## 方法

```go
func canReorderDoubled(A []int) bool {
    size := len(A)
	A1, A2 := make([]int, 0, size), make([]int, 0, size)
	for _, v := range A {
		if v >= 0 {
			A1 = append(A1, v)
		} else {
			A2 = append(A2, -v)
		}
	}
	// 把负数单独分出来，并转换成正数的数列
	// 两个部分，就可以用相同的逻辑来处理了。
	return isPossible(A1) && isPossible(A2)
}

func isPossible(A []int) bool {
	size := len(A)

	if size%2 == 1 {
		return false
	}

	sort.Ints(A)

	count := 0
	i, j := 0, 1
	for j < size {
		for A[i] < 0 {
			i++
		}
		double := A[i] * 2
		for (j < size && A[j] < double) ||
			j <= i {
			j++
		}
		if j == size || A[j] != double {
			return false
		}
		A[j] = -1
		count++
		i++
		j++
	}

	return count == size/2
}
```


```python
class Solution(object):
    def canReorderDoubled(self, A):
        """
        :type A: List[int]
        :rtype: bool
        """
        count = collections.Counter(A)
        for x in sorted(count, key=abs):
            if count[x] > count[2*x]:
                return False
            count[2*x] -= count[x]
        return True
```