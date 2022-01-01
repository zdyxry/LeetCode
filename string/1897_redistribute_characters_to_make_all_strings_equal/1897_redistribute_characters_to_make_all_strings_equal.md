1897. Redistribute Characters to Make All Strings Equal


Easy


You are given an array of strings words (0-indexed).

In one operation, pick two distinct indices i and j, where words[i] is a non-empty string, and move any character from words[i] to any position in words[j].

Return true if you can make every string in words equal using any number of operations, and false otherwise.

 

Example 1:

```
Input: words = ["abc","aabc","bc"]
Output: true
Explanation: Move the first 'a' in words[1] to the front of words[2],
to make words[1] = "abc" and words[2] = "abc".
All the strings are now equal to "abc", so return true.
```

Example 2:

```
Input: words = ["ab","a"]
Output: false
Explanation: It is impossible to make all the strings equal using the operation.
```

Constraints:

1 <= words.length <= 100   
1 <= words[i].length <= 100   
words[i] consists of lowercase English letters.   


## Python3

```
class Solution:
    def makeEqual(self, words: List[str]) -> bool:
        cnt = [0] * 26   # 每种字符的频数
        n = len(words)
        for wd in words:
            for ch in wd:
                cnt[ord(ch)-ord('a')] += 1
        return all(k % n == 0 for k in cnt)

```

## Golang

```
func makeEqual(words []string) bool {
    cnt := make(map[rune]int)
    for _, word := range(words) {
        for _, i := range(word) {
            cnt[i] += 1
        }
    }
    l := len(words)
    for _, v := range cnt {
        if v % l != 0{
            return false
        }
    }
    return true
    
}
```