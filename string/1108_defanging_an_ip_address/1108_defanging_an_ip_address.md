1108. Defanging an IP Address


Easy


Given a valid (IPv4) IP address, return a defanged version of that IP address.

A defanged IP address replaces every period "." with "[.]".

 

Example 1:

```
Input: address = "1.1.1.1"
Output: "1[.]1[.]1[.]1"
```

Example 2:

```
Input: address = "255.100.50.0"
Output: "255[.]100[.]50[.]0"
```
 

Constraints:

The given address is a valid IPv4 address.


## 方法


```go
func defangIPaddr(address string) string {
    var defanged []rune
    
    for _, char := range address{
        if char == '.' {
            defanged = append(defanged,'[','.',']')
        } else {
            defanged = append(defanged, char)
        }
    }
  
    return string(defanged)
}
```

```python
class Solution:
    def defangIPaddr(self, address: str) -> str:
        return address.replace('.', '[.]')
```