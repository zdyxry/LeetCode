150. Evaluate Reverse Polish Notation


Medium


Evaluate the value of an arithmetic expression in Reverse Polish Notation.

Valid operators are +, -, *, /. Each operand may be an integer or another expression.

Note:


Division between two integers should truncate toward zero.  
The given RPN expression is always valid. That means the expression would always evaluate to a result and there won't be any divide by zero operation.  


Example 1:

```
Input: ["2", "1", "+", "3", "*"]
Output: 9
Explanation: ((2 + 1) * 3) = 9
```

Example 2:

```
Input: ["4", "13", "5", "/", "+"]
Output: 6
Explanation: (4 + (13 / 5)) = 6
```

Example 3:
```

Input: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
Output: 22
Explanation: 
  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
= ((10 * (6 / (12 * -11))) + 17) + 5
= ((10 * (6 / -132)) + 17) + 5
= ((10 * 0) + 17) + 5
= (0 + 17) + 5
= 17 + 5
= 22
```

## 方法


```go
func evalRPN(tokens []string) int {
    // 用于存放数字的栈
	nums := make([]int, 0, len(tokens))
	for _, s := range tokens {
		if s == "+" ||
			s == "-" ||
			s == "*" ||
			s == "/" {
			// 遇到操作符， 数字出栈
			b, a := nums[len(nums)-1], nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			// 运算后的结果，重新入栈
			nums = append(nums, compute(a, b, s))
		} else {
			// 遇到数字，则直接入栈
			temp, _ := strconv.Atoi(s)
			nums = append(nums, temp)
		}
	}

	return nums[0]
}

// 计算
func compute(a, b int, opt string) int {
	switch opt {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return a / b
	}
}
```





```python
class Solution(object):
    def evalRPN(self, tokens):
        """
        :type tokens: List[str]
        :rtype: int
        """
        numerals, operators = [], {"+": operator.add, "-": operator.sub, "*": operator.mul, "/": operator.div}
        for token in tokens:
            if token not in operators:
                numerals.append(int(token))
            else:
                y, x = numerals.pop(), numerals.pop()
                numerals.append(int(operators[token](x * 1.0, y)))
        return numerals.pop()
```