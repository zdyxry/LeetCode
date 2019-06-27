1002. Find Common Characters

Easy

Given an array A of strings made only from lowercase letters, return a list of all characters that show up in all strings within the list (including duplicates).  For example, if a character occurs 3 times in all strings but not 4 times, you need to include that character three times in the final answer.

You may return the answer in any order.

 

Example 1:

```
Input: ["bella","label","roller"]
Output: ["e","l","l"]
```

Example 2:

```
Input: ["cool","lock","cook"]
Output: ["c","o"]
```
 

Note:

1 <= A.length <= 100

1 <= A[i].length <= 100

A[i][j] is a lowercase letter

# 方法
遍历列表，以 26个字母为外层循环，分别统计每个字母在单词中出现的次数，并在所有次数中取最小，若次数大于0，则按照出现次数将其添加到结果中。

我们也可以取 A[0] 元素中所有字母作为起始，这样外层循环次数就从 26 减为 A[0] 中字母次数。

```go
func commonChars(A []string) []string {
	res := make([]string, 0, 128)
	for i := 'a'; i <= 'z'; i++ {
		sub := string(i)
		count := strings.Count(A[0], sub)
		for j := 1; j < len(A) && count > 0; j++ {
			count = min(count, strings.Count(A[j], sub))
		}
		for count > 0 {
			res = append(res, sub)
			count--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```


```python
class Solution(object):
    def commonChars(self, A):
        """
        :type A: List[str]
        :rtype: List[str]
        """
        from collections import Counter
        common_counter = Counter(A[0])                
        for str in A[1:]:
            common_counter &= Counter(str)                                    
        return list(common_counter.elements())
```


```python
class Solution(object):
    def commonChars(self, A):
        """
        :type A: List[str]
        :rtype: List[str]
        """
        total_unique_chars = set(''.join(A[0]))
        result = []
        for char in total_unique_chars:
            char_count = []
            for i in A:
                char_count.append(i.count(char))
            result.extend([char]*min(char_count))  
        return result

```