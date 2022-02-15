2085. Count Common Words With One Occurrence


Easy

Given two string arrays words1 and words2, return the number of strings that appear exactly once in each of the two arrays.

 

Example 1:

```
Input: words1 = ["leetcode","is","amazing","as","is"], words2 = ["amazing","leetcode","is"]
Output: 2
Explanation:
- "leetcode" appears exactly once in each of the two arrays. We count this string.
- "amazing" appears exactly once in each of the two arrays. We count this string.
- "is" appears in each of the two arrays, but there are 2 occurrences of it in words1. We do not count this string.
- "as" appears once in words1, but does not appear in words2. We do not count this string.
Thus, there are 2 strings that appear exactly once in each of the two arrays.
```

Example 2:

```
Input: words1 = ["b","bb","bbb"], words2 = ["a","aa","aaa"]
Output: 0
Explanation: There are no strings that appear in each of the two arrays.
```

Example 3:

```
Input: words1 = ["a","ab"], words2 = ["a","a","a","ab"]
Output: 1
Explanation: The only string that appears exactly once in each of the two arrays is "ab".
```

Constraints:

```
1 <= words1.length, words2.length <= 1000
1 <= words1[i].length, words2[j].length <= 30
words1[i] and words2[j] consists only of lowercase English letters.
```


## 方法


```
func countWords(words1 []string, words2 []string) int {
    result := 0
    cnt1 := make(map[string]int)
    cnt2 := make(map[string]int)
    for _, w := range words1 {
        cnt1[w] += 1
    }
    for _, w := range words2 {
        cnt2[w] += 1
    }
    for k,v := range cnt1 {
        if v == 1 && cnt2[k] == 1 {
            result += 1
        }
    }
    return result
    
}
```

```
class Solution:
    def countWords(self, words1: List[str], words2: List[str]) -> int:
        freq1 = Counter(words1)   # words1 中字符串的出现次数
        freq2 = Counter(words2)   # words2 中字符串的出现次数
        res = 0   # 出现过一次的公共字符串个数
        for word1 in freq1.keys():
            # 遍历 words1 出现的字符并判断是否满足要求
            if freq1[word1] == 1 and freq2[word1] == 1:
                res += 1
        return res

```