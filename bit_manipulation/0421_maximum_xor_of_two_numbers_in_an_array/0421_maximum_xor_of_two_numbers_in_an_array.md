421. Maximum XOR of Two Numbers in an Array


Medium


Given a non-empty array of numbers, a0, a1, a2, … , an-1, where 0 ≤ ai < 231.

Find the maximum result of ai XOR aj, where 0 ≤ i, j < n.

Could you do this in O(n) runtime?

Example:

```
Input: [3, 10, 5, 25, 2, 8]

Output: 28

Explanation: The maximum result is 5 ^ 25 = 28.
```

## 方法

```go
func findMaximumXOR(nums []int) int {
    var max, mask int
    
    for i := 31; i >= 0; i-- {
		mask |= 1 << uint(i)

		nMap := make(map[int]bool)
		for _, num := range nums {
			nMap[num&mask] = true
		}

		tmp := max | (1 << uint32(i))
		for key := range nMap {
			if nMap[tmp^key] {
				max = tmp
				break
			}
		}
	}

	return max

}
```



```python
class Solution(object):
    def findMaximumXOR(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        result = 0

        for i in reversed(xrange(32)):
            result <<= 1
            prefixes = set()
            for n in nums:
                prefixes.add(n >> i)
            for p in prefixes:
                if (result | 1) ^ p in prefixes:
                    result += 1
                    break

        return result
```