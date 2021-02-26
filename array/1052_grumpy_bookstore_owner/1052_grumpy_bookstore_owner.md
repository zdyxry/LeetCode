1052. Grumpy Bookstore Owner


Medium

Today, the bookstore owner has a store open for customers.length minutes.  Every minute, some number of customers (customers[i]) enter the store, and all those customers leave after the end of that minute.

On some minutes, the bookstore owner is grumpy.  If the bookstore owner is grumpy on the i-th minute, grumpy[i] = 1, otherwise grumpy[i] = 0.  When the bookstore owner is grumpy, the customers of that minute are not satisfied, otherwise they are satisfied.

The bookstore owner knows a secret technique to keep themselves not grumpy for X minutes straight, but can only use it once.

Return the maximum number of customers that can be satisfied throughout the day.

 

Example 1:

```
Input: customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], X = 3
Output: 16
Explanation: The bookstore owner keeps themselves not grumpy for the last 3 minutes. 
The maximum number of customers that can be satisfied = 1 + 1 + 1 + 1 + 7 + 5 = 16.
```

Note:

1 <= X <= customers.length == grumpy.length <= 20000   
0 <= customers[i] <= 1000   
0 <= grumpy[i] <= 1   





## 方法

```go
func maxSatisfied(customers []int, grumpy []int, X int) int {
    res := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			res += customers[i]
			customers[i] = 0
		}
	}
	max_count := 0
	val := 0
	for i := 0; i < len(customers); i++ {
		if i < X{
			val += customers[i]
			continue
		}
		max_count = max(val, max_count)
		val += customers[i]
		val -= customers[i-X]
	}
	max_count = max(val, max_count)
	return res + max_count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```



```python
class Solution:
    def maxSatisfied(self, customers: List[int], grumpy: List[int], X: int) -> int:
        res = 0
        for i in range(len(customers)):
            if grumpy[i] == 0:
                res += customers[i]
                customers[i] = 0
        max_count = 0
        val = 0
        for i in range(len(customers)):
            if i < X:
                val += customers[i]
                continue
            max_count = max(val, max_count)
            val -= customers[i - X]
            val += customers[i]
        max_count = max(val, max_count)

        return res + max_count
```