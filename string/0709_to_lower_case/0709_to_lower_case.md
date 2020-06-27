709. To Lower Case


Easy


Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.

 

Example 1:

```
Input: "Hello"
Output: "hello"
```

Example 2:

```
Input: "here"
Output: "here"
```

Example 3:

```
Input: "LOVELY"
Output: "lovely"
```


## 方法

```go
func toLowerCase(str string) string {
    runes := [] rune(str)
    diff := 'a' - 'A'
    for i := 0; i < len(str); i++ {
        if runes[i] >= 'A' && runes[i] <= 'Z' {
            runes[i] += diff
        }
    }
    return string(runes)
}
```


```python
class Solution:
    def toLowerCase(self, str: str) -> str:
        return str.lower()
```