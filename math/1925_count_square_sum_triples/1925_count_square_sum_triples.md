1925. Count Square Sum Triples


Easy


A square triple (a,b,c) is a triple where a, b, and c are integers and a2 + b2 = c2.

Given an integer n, return the number of square triples such that 1 <= a, b, c <= n.

 

Example 1:

```
Input: n = 5
Output: 2
Explanation: The square triples are (3,4,5) and (4,3,5).
```

Example 2:

```
Input: n = 10
Output: 4
Explanation: The square triples are (3,4,5), (4,3,5), (6,8,10), and (8,6,10).
```

Constraints:

1 <= n <= 250


## 方法


```
class Solution:
    def countTriples(self, n: int) -> int:
        res = 0
        for a in range(1, n + 1):
            for b in range(1, n + 1):
                c = sqrt(a ** 2 + b ** 2)
                if not (c % 1) and int(c) <= n and c ** 2 == a ** 2 + b ** 2:
                    res += 1
        return res

```

```
func countTriples(n int) int {
	num := 0
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			for c := 1; c <= n; c++ {
				if c*c == a*a+b*b {
					num += 1
				}
			}
		}
	}
	return num
}

```