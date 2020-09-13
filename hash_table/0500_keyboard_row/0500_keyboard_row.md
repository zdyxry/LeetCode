500. Keyboard Row


Easy


Given a List of words, return the words that can be typed using letters of alphabet on only one row's of American keyboard like the image below.

 



 
Example:

```
Input: ["Hello", "Alaska", "Dad", "Peace"]
Output: ["Alaska", "Dad"]
```

Note:

You may use one character in the keyboard more than once.  
You may assume the input string will only contain letters of alphabet.

## 方法


```go
func findWords(words []string) []string {
    indexes := [26]int{}
    for _, v := range "qwertyuiop" {
        indexes[v-'a'] = 1
    }
    
    for _, v := range "asdfghjkl" {
        indexes[v-'a'] = 2
    }
    
    for _, v := range "zxcvbnm" {
        indexes[v-'a'] = 3
    }
    
    res := make([]string, 0)
    for _, word := range words {
        temp := strings.ToLower(word)
        found := true
        for _, c := range temp {
            if indexes[c-'a'] != indexes[temp[0]-'a'] {
                found = false
                break
            }
        }
        
        if found {
            res = append(res, word)
        }
    }
    
    return res
}

```


```python
class Solution:
    def findWords(self, words: List[str]) -> List[str]:
        set1 = set('qwertyuiop')
        set2 = set('asdfghjkl')
        set3 = set('zxcvbnm')
        res = []
        for i in words:
            x = i.lower()
            setx = set(x)
            if setx<=set1 or setx<=set2 or setx<=set3:
                res.append(i)
        return res
```