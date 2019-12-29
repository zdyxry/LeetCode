405. Convert a Number to Hexadecimal


Easy


Given an integer, write an algorithm to convert it to hexadecimal. For negative integer, two’s complement method is used.

Note:

All letters in hexadecimal (a-f) must be in lowercase.  
The hexadecimal string must not contain extra leading 0s. If the number is zero, it is represented by a single zero character '0'; otherwise, the first character in the hexadecimal string will not be the zero character.  
The given number is guaranteed to fit within the range of a 32-bit signed integer.  
You must not use any method provided by the library which converts/formats the number to hex directly.  

Example 1:

```
Input:
26

Output:
"1a"
```

Example 2:

```
Input:
-1

Output:
"ffffffff"
```


## 方法

```go
var h = []string{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"a",
	"b",
	"c",
	"d",
	"e",
	"f",
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	hex := ""
	for i := 0; i < 8 && num != 0; i++ {
		hex = h[num&15] + hex
		num >>= 4
	}

	return hex
}
```


```python
class Solution(object):
    def toHex(self, num):
        """
        :type num: int
        :rtype: str
        """
        if not num:
            return "0"

        result = []
        while num and len(result) != 8:
            h = num & 15
            if h < 10:
                result.append(str(chr(ord('0') + h)))
            else:
                result.append(str(chr(ord('a') + h-10)))
            num >>= 4
        result.reverse()

        return "".join(result)
```