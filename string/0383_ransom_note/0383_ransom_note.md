383. Ransom Note


Easy


Given an arbitrary ransom note string and another string containing letters from all the magazines, write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.

Note:
You may assume that both strings contain only lowercase letters.

```
canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true
```


## 方法

使用map 或 list 保存 magazine 字母出现次数，然后遍历 ransomNote。

```go
func canConstruct(ransomNote string, magazine string) bool {
    mc := getCount(magazine)

	for _, b := range ransomNote {
		mc[b-'a']--
		if mc[b-'a'] < 0 {
			return false
		}
	}

	return true
}

func getCount(s string) []int {
	res := make([]int, 26)
	for i := range s {
		res[s[i]-'a']++
	}
	return res
}

```



```python
class Solution(object):
    def canConstruct(self, ransomNote, magazine):
        """
        :type ransomNote: str
        :type magazine: str
        :rtype: bool
        """
        dic = {}
        for s in ransomNote:
            dic[s] = ransomNote.count(s)
        for key in dic.keys():
            if magazine.count(key) < dic[key]:
                return False
        return True

```


```python
class Solution(object):
    def canConstruct(self, ransomNote, magazine):
        """
        :type ransomNote: str
        :type magazine: str
        :rtype: bool
        """
        counts = [0] * 26
        letters = 0

        for c in ransomNote:
            if counts[ord(c) - ord('a')] == 0:
                letters += 1
            counts[ord(c) - ord('a')] += 1

        for c in magazine:
            counts[ord(c) - ord('a')] -= 1
            if counts[ord(c) - ord('a')] == 0:
                letters -= 1
                if letters == 0:
                    break

        return letters == 0
```