925. Long Pressed Name


Easy


Your friend is typing his name into a keyboard.  Sometimes, when typing a character c, the key might get long pressed, and the character will be typed 1 or more times.

You examine the typed characters of the keyboard.  Return True if it is possible that it was your friends name, with some characters (possibly none) being long pressed.

 

Example 1:
```
Input: name = "alex", typed = "aaleex"
Output: true
Explanation: 'a' and 'e' in 'alex' were long pressed.
```

Example 2:
```
Input: name = "saeed", typed = "ssaaedd"
Output: false
Explanation: 'e' must have been pressed twice, but it wasn't in the typed output.
```

Example 3:
```
Input: name = "leelee", typed = "lleeelee"
Output: true
```

Example 4:
```
Input: name = "laiden", typed = "laiden"
Output: true
Explanation: It's not necessary to long press any character.
```

Note:

name.length <= 1000

typed.length <= 1000

The characters of name and typed are lowercase letters.



## 方法

双指针，比较分别遍历两个字符串，当遇到 name[i] == typed[j] 时，i与j 都 +1，当遇到不等时，j +1 。


```go
func isLongPressedName(name string, typed string) bool {
    if name == typed {
		return true
	}

	nameSize := len(name)
	typedSize := len(typed)

	i, j := 0, 0

	for i < nameSize && j < typedSize {
		c := name[i]
		need, pressed := 0, 0

		for i < nameSize && name[i] == c {
			need++
			i++
		}

		for j < typedSize && typed[j] == c {
			pressed++
			j++
		}

		if pressed < need {
			return false
		}

	}

	return i == nameSize && j == typedSize
}
```





```python
class Solution(object):
    def isLongPressedName(self, name, typed):
        """
        :type name: str
        :type typed: str
        :rtype: bool
        """
        i = 0
        for j in xrange(len(typed)):
            if i < len(name) and name[i] == typed[j]:
                i += 1
            elif j == 0 or typed[j] != typed[j-1]:
                return False
        return i == len(name)
```