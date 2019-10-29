451. Sort Characters By Frequency


Medium


Given a string, sort it in decreasing order based on the frequency of characters.

Example 1:

```
Input:
"tree"

Output:
"eert"

Explanation:
'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
```

Example 2:

```
Input:
"cccaaa"

Output:
"cccaaa"

Explanation:
Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
Note that "cacaca" is incorrect, as the same characters must be together.
```

Example 3:

```
Input:
"Aabb"

Output:
"bbAa"

Explanation:
"bbaA" is also a valid answer, but "Aabb" is incorrect.
Note that 'A' and 'a' are treated as two different characters.
```

## 方法

```go
func frequencySort(s string) string {
    if s == "" {
		return ""
	}
	sMap := map[byte]int{}
	cMap := map[int][]byte{}
	sb := []byte(s)
	for _, b := range sb {
		sMap[b]++
	}
	for key, value := range sMap {
		cMap[value] = append(cMap[value], key)
	}

	var keys []int
	for k := range cMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	res := make([]byte, 0)
	for _, k := range keys {
		for i := 0; i < len(cMap[k]); i++ {
			for j := 0; j < k; j++ {
				res = append(res, cMap[k][i])
			}
		}
	}
	return string(res)
}
```


```python
class Solution(object):
    def frequencySort(self, s):
        """
        :type s: str
        :rtype: str
        """
        # dictionary to store count of character appearing in the string
        dic = {}
		
		# final result
        result = ""
		
		# loop through each char, and add the count to the dic
        for char in s:
            if char in dic:
                dic[char] += 1
            else:
                dic[char] = 1
				
		# sort the dic on values in reverse order
        sorted_dic = sorted(dic, key=dic.get, reverse=True)
		
		# loop through, and create final string
        for count in sorted_dic:
            result += count * (dic[count])

        return result
```