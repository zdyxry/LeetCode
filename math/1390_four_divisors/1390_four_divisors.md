1390. Four Divisors


Medium


Given an integer array nums, return the sum of divisors of the integers in that array that have exactly four divisors.

If there is no such integer in the array, return 0.

 

Example 1:

```
Input: nums = [21,4,7]
Output: 32
Explanation:
21 has 4 divisors: 1, 3, 7, 21
4 has 3 divisors: 1, 2, 4
7 has 2 divisors: 1, 7
The answer is the sum of divisors of 21 only.
```

Constraints:

1 <= nums.length <= 10^4   
1 <= nums[i] <= 10^5


## 方法

```go
func sumFourDivisors(nums []int) int {
    summary := 0
	for _, v := range nums {
		summary += helper(v)
	}
	return summary
}

func helper(num int) int {
	if num <= 5 {
		return 0
	}
	root := int(math.Sqrt(float64(num)))
    if root * root == num {
        return 0
    }
	divisorCount := 2
    summary := num + 1
	for i := 2; i <= root; i++ {
		if num % i == 0 {
			divisorCount += 2
            summary += i
            summary += num / i
			if divisorCount > 4 {
				return 0
			}
		}
	}
    if divisorCount == 4 {
        return summary
    }
    return 0
}
```



```python
class Solution(object):
    def sumFourDivisors(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        cache = {}
        def fn(n):
            if n not in cache:
                s = set()
                for i in range(1, 1 + int(math.sqrt(n))):
                    if n % i == 0:
                        s.add(i)
                        s.add(n / i)
                    if len(s) > 4:
                        break
                cache[n] = sum(s) if len(s) == 4 else 0
            return cache[n]
                
        return sum(map(fn, nums))
```