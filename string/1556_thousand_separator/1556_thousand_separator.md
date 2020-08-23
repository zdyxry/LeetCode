1556. Thousand Separator


Easy


Given an integer n, add a dot (".") as the thousands separator and return it in string format.

 

Example 1:

```
Input: n = 987
Output: "987"
```

Example 2:

```
Input: n = 1234
Output: "1.234"
```

Example 3:

```
Input: n = 123456789
Output: "123.456.789"
```

Example 4:

```
Input: n = 0
Output: "0"
```
 

Constraints:

0 <= n < 2^31


## 方法

```go
func reverse(s string) string {
    temp:=[]rune(s)
    i:=0
    j:=len(temp)-1
    for i<j {
        temp[i],temp[j]=temp[j],temp[i]
        i++
        j--
    }
    return string(temp)
}

func thousandSeparator(n int) string {
    rev_s:=reverse(strconv.Itoa(n))
    pattern:=regexp.MustCompile(`\d{1,3}`)
    return reverse(strings.Join(pattern.FindAllString(rev_s,-1),"."))
}
```

```python
class Solution:
    def thousandSeparator(self, n: int) -> str:
        n = str(n)[::-1]
        n1 = ""
        c = 0
        for i in n:
            c += 1
            n1 += i
            if c % 3 == 0:
                n1 += "."
        return n1[::-1].strip(".")
```