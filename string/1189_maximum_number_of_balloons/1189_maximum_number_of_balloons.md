1189. Maximum Number of Balloons


Easy


Given a string text, you want to use the characters of text to form as many instances of the word "balloon" as possible.

You can use each character in text at most once. Return the maximum number of instances that can be formed.

 

Example 1:



```
Input: text = "nlaebolko"
Output: 1
```

Example 2:



```
Input: text = "loonbalxballpoon"
Output: 2
```

Example 3:

```
Input: text = "leetcode"
Output: 0
```
 

Constraints:

1 <= text.length <= 10^4  
text consists of lower case English letters only.

## 方法

```go
func maxNumberOfBalloons(text string) int {
    l :=[26]int{}
	for _,b :=range text{
		l[b-'a']++
	}
	l['l'-'a']= l['l'-'a']/2
	l['o'-'a']= l['o'-'a']/2
	var min = math.MaxInt16
	for _,b :=range "balon"{
		  v :=l[b-'a']
          if v==0{
          	return 0
		  }

		if v<min {
			min = v
		}
	}
	return min
}
```

```python
class Solution:
    def maxNumberOfBalloons(self, text: str) -> int:
        d = collections.Counter(text)
        return min(d['b'], d['a'], d['l'] // 2, d['o'] // 2, d['n'])
```