1048. Longest String Chain


Medium


Given a list of words, each word consists of English lowercase letters.

Let's say word1 is a predecessor of word2 if and only if we can add exactly one letter anywhere in word1 to make it equal to word2.  For example, "abc" is a predecessor of "abac".

A word chain is a sequence of words [word_1, word_2, ..., word_k] with k >= 1, where word_1 is a predecessor of word_2, word_2 is a predecessor of word_3, and so on.

Return the longest possible length of a word chain with words chosen from the given list of words.

 

Example 1:

```
Input: ["a","b","ba","bca","bda","bdca"]
Output: 4
Explanation: one of the longest word chain is "a","ba","bda","bdca".
```

Note:

1 <= words.length <= 1000  
1 <= words[i].length <= 16  
words[i] only consists of English lowercase letters.


## 方法

```go
func longestStrChain(words []string) int {
	if len(words) < 2 {
		return len(words)
	}

	var dp = make([]int,len(words))
	var res = 0
    dp[0] = 1 
	//排序字符串数组
	sort.Slice(words,func(i,j int)bool{
		if len(words[i]) < len(words[j]){
			return true
		}else{
			return false
		}
	})
	//核心代码
	for i:=1;i<len(words);i++{
		dp[i] = 1 //初始化每一位置的值为1
		for j := 0;j < i;j++{
			//检查是否是子串
			if isChain(words[j],words[i]){
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
		res = int(math.Max(float64(res), float64(dp[i])))
	}
	return res
}
//检查子串
func isChain(a string,b string) bool{
	//必备条件：长度相差1
	if len(b) - len(a) == 1{
		var i = 0
		var j = 0
		//勉强算是双指针吧
		for i <len(a) && j < len(b) {
			if a[i] == b[j] {
				i++
			}
			j++
		}
		if i == len(a){
			return true
		}
	}
	return false
}

```



```python
class Solution:
    def longestStrChain(self, words: List[str]) -> int:
        words.sort(key=len)
        note={}
        maxChain=1
        for word in words:
            if word not in note:
                note[word]=1
            for i in range(0,len(word)):
                newWord=word[:i]+word[i+1:]
                if (newWord) in note:
                    note[word]=max(note[word],note[newWord]+1)
            maxChain=max(maxChain,note[word])
        return maxChain
```