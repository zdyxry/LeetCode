1362. Closest Divisors


Medium


Given an integer num, find the closest two integers in absolute difference whose product equals num + 1 or num + 2.

Return the two integers in any order.

 

Example 1:

```
Input: num = 8
Output: [3,3]
Explanation: For num + 1 = 9, the closest divisors are 3 & 3, for num + 2 = 10, the closest divisors are 2 & 5, hence 3 & 3 is chosen.
```

Example 2:

```
Input: num = 123
Output: [5,25]
```

Example 3:

```
Input: num = 999
Output: [40,25]
```

Constraints:

1 <= num <= 10^9


## 方法


```go
func closestDivisors(num int) []int {
    v1 := closestDivisorsHelper(num+1)
	v2 := closestDivisorsHelper(num+2)
	if v1[1] - v1[0] < v2[1] - v2[0] {
		return v1
	}
	return v2
}

func closestDivisorsHelper(num int) []int {
	root := int(math.Sqrt(float64(num)))
	for i := root; i >= 1; i-- {
		if num % i == 0 {
			return []int{i, num/i}
		}
	}
	return []int{}
}
```


```python
class Solution(object):
    def closestDivisors(self, num):
        """
        :type num: int
        :rtype: List[int]
        """
        for i in reversed(range(1, int((num + 2) ** 0.5) + 1)):
            if not (num + 1) % i:
                return [i, (num + 1) // i]
            if not (num + 2) % i:
                return [i, (num + 2) // i]
```