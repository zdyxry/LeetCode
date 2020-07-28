1523. Count Odd Numbers in an Interval Range

Easy

Given two non-negative integers low and high. Return the count of odd numbers between low and high (inclusive).

 

Example 1:

```
Input: low = 3, high = 7
Output: 3
Explanation: The odd numbers between 3 and 7 are [3,5,7].
```

Example 2:

```
Input: low = 8, high = 10
Output: 1
Explanation: The odd numbers between 8 and 10 are [9].
```

 

Constraints:

0 <= low <= high <= 10^9

## 方法

```go
func countOdds(low int, high int) int {
    if low +1 == high {
        return 1
    }
    if low == high {
        if low % 2 == 0 {
            return 0
        }
        return 1
    }
    if low % 2 == 0 {
        low++
    }
    num := (high - low) / 2
    return num +1
}

```



```python
class Solution:
    def countOdds(self, low: int, high: int) -> int:
        res = high // 2
        if high % 2 == 1:
            res += 1
        if low - 1 >= 0:
            res -= (low - 1) // 2
            if (low - 1) % 2 == 1:
                res -= 1
        return res
```
