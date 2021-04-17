1805. Number of Different Integers in a String


Easy


You are given a string word that consists of digits and lowercase English letters.

You will replace every non-digit character with a space. For example, "a123bc34d8ef34" will become " 123  34 8  34". Notice that you are left with some integers that are separated by at least one space: "123", "34", "8", and "34".

Return the number of different integers after performing the replacement operations on word.

Two integers are considered different if their decimal representations without any leading zeros are different.

 

Example 1:

```
Input: word = "a123bc34d8ef34"
Output: 3
Explanation: The three different integers are "123", "34", and "8". Notice that "34" is only counted once.
```

Example 2:

```
Input: word = "leet1234code234"
Output: 2
```

Example 3:

```
Input: word = "a1b01c001"
Output: 1
Explanation: The three integers "1", "01", and "001" all represent the same integer because
the leading zeros are ignored when comparing their decimal values.
```

Constraints:

1 <= word.length <= 1000   
word consists of digits and lowercase English letters.   


## 方法

```go
func numDifferentIntegers(word string) int {
    all := make(map[string]bool)
    
    for i := 0; i < len(word); i++ {
        c := word[i]
        // ignore leading zeros
        if c == '0' && i + 1 < len(word) && isDigit(word[i + 1]) {
            continue
        }
        if !isDigit(c) {
            continue
        }
        // when first digit encountered, let's collect the slice to the map
        start := i
        i++
        for i < len(word) && isDigit(word[i]) {
            i++
        }
        all[word[start:i]] = true
    }
    
    return len(all)
}

func isDigit(c byte) bool {    
    return '0' <= c && c <= '9';
}
```



```python
class Solution:
    def numDifferentIntegers(self, word: str) -> int:
        ret = set()
        stack = []
        for i in word:
            if i.isdigit():
                if len(stack) == 1 and stack[-1] == '0':
                    stack.pop()
                stack.append(i)
            else:
                if stack:
                    ret.add(''.join(stack))
                stack = []
        if stack:
            ret.add(''.join(stack))
        return len(ret)
```