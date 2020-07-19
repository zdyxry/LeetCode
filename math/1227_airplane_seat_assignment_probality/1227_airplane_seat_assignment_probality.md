1227. Airplane Seat Assignment Probability


Medium


n passengers board an airplane with exactly n seats. The first passenger has lost the ticket and picks a seat randomly. But after that, the rest of passengers will:


Take their own seat if it is still available,   
Pick other seats randomly when they find their seat occupied   
What is the probability that the n-th person can get his own seat?

 

Example 1:

```
Input: n = 1
Output: 1.00000
Explanation: The first person can only get the first seat.
```

Example 2:

```
Input: n = 2
Output: 0.50000
Explanation: The second person has a probability of 0.5 to get the second seat (when first person gets the first seat).
```
 

Constraints:

1 <= n <= 10^5


## 方法

```
func nthPersonGetsNthSeat(n int) float64 {
    if n == 1 {
        return 1
    }
    return 0.5
}
```



```python
class Solution:
    def nthPersonGetsNthSeat(self, n: int) -> float:
        return 1 if n==1 else .5
```