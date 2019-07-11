917. Reverse Only Letters


Easy


Given a string S, return the "reversed" string where all characters that are not a letter stay in the same place, and all letters reverse their positions.

 

Example 1:
```
Input: "ab-cd"
Output: "dc-ba"
```
Example 2:
```
Input: "a-bC-dEf-ghIj"
Output: "j-Ih-gfE-dCba"
```
Example 3:
```
Input: "Test1ng-Leet=code-Q!"
Output: "Qedo1ct-eeLg=ntse-T!"
```


## 方法
反转字符串，判断字符是否为字母，若为字母，则交换，否则跳过该字符。


```go
func reverseOnlyLetters(S string) string {
    bs := []byte(S)

	left, right := 0, len(bs)-1
	for left < right {
		for left < right && !isLetter(bs[left]) {
			left++
		}
		for left < right && !isLetter(bs[right]) {
			right--
		}
		bs[left], bs[right] = bs[right], bs[left]
		left++
		right--
	}

	return string(bs)
}

func isLetter(b byte) bool {
	return 'a' <= b && b <= 'z' ||
		'A' <= b && b <= 'Z'
}

```


```python
class Solution(object):
    def reverseOnlyLetters(self, S):
        """
        :type S: str
        :rtype: str
        """
        def getNext(S):
            for i in reversed(xrange(len(S))):
                if S[i].isalpha():
                    yield S[i]

        result = []
        letter = getNext(S)
        for i in xrange(len(S)):
            if S[i].isalpha():
                result.append(letter.next())
            else:
                result.append(S[i])
        return "".join(result)

```