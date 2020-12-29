1702. Maximum Binary String After Change


Medium


You are given a binary string binary consisting of only 0's or 1's. You can apply each of the following operations any number of times:

```
Operation 1: If the number contains the substring "00", you can replace it with "10".
For example, "00010" -> "10010"
Operation 2: If the number contains the substring "10", you can replace it with "01".
For example, "00010" -> "00001"
Return the maximum binary string you can obtain after any number of operations. Binary string x is greater than binary string y if x's decimal representation is greater than y's decimal representation.
```

 

Example 1:

```
Input: binary = "000110"
Output: "111011"
Explanation: A valid transformation sequence can be:
"000110" -> "000101" 
"000101" -> "100101" 
"100101" -> "110101" 
"110101" -> "110011" 
"110011" -> "111011"
```

Example 2:

```
Input: binary = "01"
Output: "01"
Explanation: "01" cannot be transformed any further.
```

Constraints:

1 <= binary.length <= 105   
binary consist of '0' and '1'.


## 方法


```go
func maximumBinaryString(binary string) string {

    firstZero := strings.IndexByte(binary, '0')
    if firstZero == -1 || firstZero == len(binary)-1 {
        return binary
    }
    res := []byte(binary)
    ones := strings.Count(binary[firstZero:], "1")
    for i := range res {
        res[i] = '1'
    }
    res[len(res)-1-ones] = '0'
    return string(res)
}
```



```python
class Solution:
    def maximumBinaryString(self, binary: str) -> str:
        n = len(binary)
        cnt0 = 0
        start = -1
        for i in range(n):
            if binary[i] == '0':
                cnt0 += 1
                if start == -1:
                    start = i
        if start == -1: return binary
        return '1' * (start + cnt0 - 1) + '0' + '1' * (n - start - cnt0)
```