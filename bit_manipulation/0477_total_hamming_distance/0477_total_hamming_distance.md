477. Total Hamming Distance


Medium


The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

Now your job is to find the total Hamming distance between all pairs of the given numbers.

Example:
```
Input: 4, 14, 2

Output: 6
```


Explanation: In binary representation, the 4 is 0100, 14 is 1110, and 2 is 0010 (just
showing the four bits relevant in this case). So the answer will be:

HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.

Note:

Elements of the given array are in the range of 0 to 10^9

Length of the array will not exceed 10^4.


## 方法


```go
func totalHammingDistance(nums []int) int {
    ans, numsLen := 0, len(nums)
	for i := 0; i < 32; i++ {
		haveOne := 0
		for _, num := range nums {
			haveOne += (num >> i) & 1  // 判断此位是否为 1
		}
		ans += haveOne * (numsLen - haveOne)
	}
	return ans

}
```


```python
class Solution:
    def totalHammingDistance(self, nums: List[int]) -> int:
        res, n = 0, len(nums)
        for i in range(32):
            cnt_1 = 0
            for j in range(n):
                cnt_1 += (nums[j] >> i) & 1
            res += (n - cnt_1) * cnt_1
        return res 

```