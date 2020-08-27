1079. Letter Tile Possibilities


Medium


You have n  tiles, where each tile has one letter tiles[i] printed on it.

Return the number of possible non-empty sequences of letters you can make using the letters printed on those tiles.

 

Example 1:

```
Input: tiles = "AAB"
Output: 8
Explanation: The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".
```

Example 2:

```
Input: tiles = "AAABBC"
Output: 188
```

Example 3:

```
Input: tiles = "V"
Output: 1
```
 

Constraints:

1 <= tiles.length <= 7  
tiles consists of uppercase English letters.


## 方法


```go
func numTilePossibilities(tiles string) int {
	used := make([]bool, len(tiles))
	m := make(map[string]struct{})
	backtrack("", used, tiles, m)
	return len(m)
}

func backtrack(now string, used []bool, str string, m map[string]struct{}) {
	if len(now) > 0 {
		m[now] = struct{}{}
	}
	if len(now) == len(str) {
		return
	}
	for i := 0; i < len(str); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		backtrack(now + string(str[i]), used, str, m)
		used[i] = false
	}
}

```


```python
class Solution:
    def numTilePossibilities(self, tiles: str) -> int:
        record = [0] * 26
        for tile in tiles: record[ord(tile)-ord('A')] += 1
        def dfs(record):
            s = 0
            for i in range(26):
                if not record[i]: continue
                record[i] -= 1
                s += dfs(record) + 1 
                record[i] += 1
            return s    
        return dfs(record)
```