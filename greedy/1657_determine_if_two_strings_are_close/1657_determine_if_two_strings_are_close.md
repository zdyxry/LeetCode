1657. Determine if Two Strings Are Close


Medium


Two strings are considered close if you can attain one from the other using the following operations:

```
Operation 1: Swap any two existing characters.
For example, abcde -> aecdb
Operation 2: Transform every occurrence of one existing character into another existing character, and do the same with the other character.
For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
You can use the operations on either string as many times as necessary.
```

Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.

 

Example 1:

```
Input: word1 = "abc", word2 = "bca"
Output: true
Explanation: You can attain word2 from word1 in 2 operations.
Apply Operation 1: "abc" -> "acb"
Apply Operation 1: "acb" -> "bca"
```

Example 2:

```
Input: word1 = "a", word2 = "aa"
Output: false
Explanation: It is impossible to attain word2 from word1, or vice versa, in any number of operations.
```

Example 3:

```
Input: word1 = "cabbba", word2 = "abbccc"
Output: true
Explanation: You can attain word2 from word1 in 3 operations.
Apply Operation 1: "cabbba" -> "caabbb"
Apply Operation 2: "caabbb" -> "baaccc"
Apply Operation 2: "baaccc" -> "abbccc"
```

Example 4:

```
Input: word1 = "cabbba", word2 = "aabbss"
Output: false
Explanation: It is impossible to attain word2 from word1, or vice versa, in any amount of operations.
```
 

Constraints:

1 <= word1.length, word2.length <= 105   
word1 and word2 contain only lowercase English letters.


## 方法

```go
func closeStrings(word1 string, word2 string) bool {
    if len(word1)  != len(word2){
        return false
    }
    
    maps := make(map[rune]int)

    for _,val:= range word1{
        maps[val] += 1
    }
    for _,val:= range word2{
        if maps[val] == 0{
            return false           
        }
        maps[val +'a'] +=1
    }


    s := 0

    for k,v:= range maps{
        if k <= 'z' {
            if maps[k +'a'] == 0{
                return false
            }
        }
        s ^= v
    }
    return s == 0

}

```


```python
class Solution:
    def closeStrings(self, word1: str, word2: str) -> bool:
        c1 = collections.Counter(word1)
        c2 = collections.Counter(word2)
        
        l1 = list(c1.values())
        l2 = list(c2.values())
        
        s1 = set(word1)
        s2 = set(word2)
        
        l1.sort()
        l2.sort()
        
        return l1 == l2 and s1 == s2
```