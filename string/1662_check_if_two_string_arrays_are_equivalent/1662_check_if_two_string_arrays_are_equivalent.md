1662. Check If Two String Arrays are Equivalent


Easy


Given two string arrays word1 and word2, return true if the two arrays represent the same string, and false otherwise.

A string is represented by an array if the array elements concatenated in order forms the string.

 

Example 1:

```
Input: word1 = ["ab", "c"], word2 = ["a", "bc"]
Output: true
Explanation:
word1 represents string "ab" + "c" -> "abc"
word2 represents string "a" + "bc" -> "abc"
The strings are the same, so return true.
```

Example 2:

```
Input: word1 = ["a", "cb"], word2 = ["ab", "c"]
Output: false
```

Example 3:

```
Input: word1  = ["abc", "d", "defg"], word2 = ["abcddefg"]
Output: true
```
 

Constraints:

1 <= word1.length, word2.length <= 103   
1 <= word1[i].length, word2[i].length <= 103   
1 <= sum(word1[i].length), sum(word2[i].length) <= 103   
word1[i] and word2[i] consist of lowercase letters.


## 方法



```go
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    if len(word1) == 0 && len(word2) == 0 {
        return true
    }
    
    full1 := strings.Join(word1, "")
    full2 := strings.Join(word2, "")
    if len(full1) != len(full2) {
        return false
    }
    
    for i := 0; i < len(full1); i++ {
        if full1[i] != full2[i] {
            return false
        }
    }
    return true
}

```


```python
class Solution:
    def arrayStringsAreEqual(self, word1: List[str], word2: List[str]) -> bool:
        return ''.join(word1) == ''.join(word2)
```