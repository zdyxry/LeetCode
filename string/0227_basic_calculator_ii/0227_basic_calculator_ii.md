227. Basic Calculator II


Medium


Implement a basic calculator to evaluate a simple expression string.

The expression string contains only non-negative integers, +, -, *, / operators and empty spaces . The integer division should truncate toward zero.

Example 1:

```
Input: "3+2*2"
Output: 7
```

Example 2:

```
Input: " 3/2 "
Output: 1
```

Example 3:

```
Input: " 3+5 / 2 "
Output: 5
```

Note:

You may assume that the given expression is always valid.  
Do not use the eval built-in library function.


## 方法

```go
func calculate(s string) int {
    	// n1 opt1 n2 opt2 n3
	// ↓   ↓   ↓   ↓   ↓
	// 1   +   2   *   3
	var n1, n2, n3 int
	var opt1, opt2 byte
	opt1 = '+'

	idx := 0
	nextN := func() int {
		n := 0
		// 跳过空格
		for idx < len(s) && s[idx] == ' ' {
			idx++
		}
		// 获取数值
		for idx < len(s) && '0' <= s[idx] && s[idx] <= '9' {
			n = n*10 + int(s[idx]-'0')
			idx++
		}
		return n
	}

	nextOpt := func() byte {
		// opt 默认为 +，与 nextN() 的默认值为 0 配合
		// 当 s 的结尾为 空格 的时候，程序会进行一次 +0 的运算，但这不会影响结果。
		opt := byte('+')
		// 跳过空格
		for idx < len(s) && s[idx] == ' ' {
			idx++
		}
		// 获取操作符
		if idx < len(s) {
			opt = s[idx]
			idx++
		}
		return opt
	}

	n2 = nextN()
	for idx < len(s) {
		opt2 = nextOpt()
		n3 = nextN()

		if opt2 == '*' || opt2 == '/' {
			// 先乘除
			n2 = operate(n2, n3, opt2)
		} else {
			// 后加减
			n1 = operate(n1, n2, opt1)
			opt1 = opt2
			n2 = n3
		}
	}

	return operate(n1, n2, opt1)
}

func operate(a, b int, opt byte) int {
	switch opt {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	default: // '/'
		return a / b
	}
}
```


```python
class Solution(object):
    def calculate(self, s):
        """
        :type s: str
        :rtype: int
        """
        if not s:
            return "0"
        stack, num, sign = [], 0, "+"
        for i in xrange(len(s)):
            if s[i].isdigit():
                num = num*10+ord(s[i])-ord("0")
            if (not s[i].isdigit() and not s[i].isspace()) or i == len(s)-1:
                if sign == "-":
                    stack.append(-num)
                elif sign == "+":
                    stack.append(num)
                elif sign == "*":
                    stack.append(stack.pop()*num)
                else:
                    tmp = stack.pop()
                    if tmp//num < 0 and tmp%num != 0:
                        stack.append(tmp//num+1)
                    else:
                        stack.append(tmp//num)
                sign = s[i]
                num = 0
        return sum(stack)
```