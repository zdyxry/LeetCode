66. Plus One

Easy

Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:
```
Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
```

Example 2:

```
Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
```

# 方法
1. 简单粗暴，将数组转换为字符串处理，处理完后转换为数组返回
2. 对原数组进行从最后位遍历，判断是否为 9，若为 9，则将其置为 1，当遍历过程中遇到不为 9，则将其 +1 返回数组，若所有位均为 9，则将在数组前加 1.


```go
func plusOne(digits []int) []int {
	length := len(digits)
	if length == 0 {
		return []int{1}
	}

	// 末尾加一
	digits[length-1]++

	// 处理进位
	for i := length - 1; i > 0; i-- {
		if digits[i] < 10 {
			break
		}
		digits[i] -= 10
		digits[i-1]++
	}

	// 处理首位的进位
	if digits[0] > 9 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}

	return digits
}

```



```python
class Solution(object):
    def plusOne(self, digits):
            """
            :type digits: List[int]
            :rtype: List[int]
            """
            for i in reversed(xrange(len(digits))):
                if digits[i] == 9:
                    digits[i] = 0
                else:
                    digits[i] += 1
                    return digits
            digits[0] = 1
            digits.append(0)
            return digits
```

```python
class Solution(object):
    def plusOne(self, digits):
        right = len(digits) - 1
        while digits[right] == 9:
            digits[right] = 0
            right -= 1
        if right < 0:
            digits = [1] + digits
        else:
            digits[right] += 1
        return digits
```