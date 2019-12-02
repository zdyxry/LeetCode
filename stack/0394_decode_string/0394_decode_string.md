394. Decode String


Medium


Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.

Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there won't be input like 3a or 2[4].

Examples:

```
s = "3[a]2[bc]", return "aaabcbc".
s = "3[a2[c]]", return "accaccacc".
s = "2[abc]3[cd]ef", return "abcabccdcdcdef".
```


## 方法


```go
func decodeString(s string) string {
    n := len(s)

	// i 是第一个数字的位置
	i := 0
	for i < n && (s[i] < '0' || '9' < s[i]) {
		i++
	}
	if i == n {
		// 没有数字，直接返回 s
		return s
	}

	// j 是第一个 '[' 的位置
	j := i + 1
	// 由题意可知，s 很规范
	// 存在数字的话，必定存在 '[' 和 ']'
	for s[j] != '[' {
		j++
	}

	// k 是与 j 的 '[' 对应的 ']' 的位置
	k := j
	count := 1
	for count > 0 {
		k++

		if s[k] == '[' {
			count++
		} else if s[k] == ']' {
			count--
		}
	}

	//      i：第一个数字的位置
	//      |  j：第一个 '[' 的位置
	//      |  |       k：与 j 的 '[' 对应的 ']' 的位置
	//      ↓  ↓       ↓
	// "abcd234[*******]efg"

	// 题目说了， s 很规范
	num, _ := strconv.Atoi(s[i:j])

	return s[:i] + times(num, decodeString(s[j+1:k])) + decodeString(s[k+1:])
}

func times(n int, s string) string {
	res := ""
	for i := 0; i < n; i++ {
		res += s
	}
	return res
}
```

```go
func decodeString(s string) string {
    stack, res := []string{}, ""
	for _, str := range s {
		if len(stack) == 0 || (len(stack) > 0 && str != ']') {
			stack = append(stack, string(str))
		} else {
			tmp := ""
			for stack[len(stack)-1] != "[" {
				tmp = stack[len(stack)-1] + tmp
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			index, repeat := 0, ""
			for index = len(stack) - 1; index >= 0; index-- {
				if stack[index] >= "0" && stack[index] <= "9" {
					repeat = stack[index] + repeat
				} else {
					break
				}
			}
			nums, _ := strconv.Atoi(repeat)
			copyTmp := tmp
			for i := 0; i < nums-1; i++ {
				tmp += copyTmp
			}
			for i := 0; i < len(repeat)-1; i++ {
				stack = stack[:len(stack)-1]
			}
			stack[index+1] = tmp
		}
	}
	for _, s := range stack {
		res += s
	}
	return res
}
```

```python
class Solution(object):
    def decodeString(self, s):
        """
        :type s: str
        :rtype: str
        """
        stack = []; curNum = 0; curString = ''
        for c in s:
            if c == '[':
                stack.append(curString)
                stack.append(curNum)
                curString = ''
                curNum = 0
            elif c == ']':
                num = stack.pop()
                prevString = stack.pop()
                curString = prevString + num*curString
            elif c.isdigit():
                curNum = curNum*10 + int(c)
            else:
                curString += c
        return curString
```