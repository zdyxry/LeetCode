2007. Find Original Array From Doubled Array


Medium


An integer array original is transformed into a doubled array changed by appending twice the value of every element in original, and then randomly shuffling the resulting array.

Given an array changed, return original if changed is a doubled array. If changed is not a doubled array, return an empty array. The elements in original may be returned in any order.

 

Example 1:

```
Input: changed = [1,3,4,2,6,8]
Output: [1,3,4]
Explanation: One possible original array could be [1,3,4]:
- Twice the value of 1 is 1 * 2 = 2.
- Twice the value of 3 is 3 * 2 = 6.
- Twice the value of 4 is 4 * 2 = 8.
Other original arrays could be [4,3,1] or [3,1,4].
```

Example 2:

```
Input: changed = [6,3,0,1]
Output: []
Explanation: changed is not a doubled array.
```

Example 3:

```
Input: changed = [1]
Output: []
Explanation: changed is not a doubled array.
```

Constraints:

1 <= changed.length <= 105   
0 <= changed[i] <= 105


## 方法


```
func findOriginalArray(changed []int) (original []int) {
	sort.Ints(changed)
	cnt := map[int]int{}
	for _, v := range changed {
		if cnt[v] == 0 { 
			cnt[v*2]++ 
			original = append(original, v)
		} else {
			cnt[v]-- 
			if cnt[v] == 0 {
				delete(cnt, v)
			}
		}
	}
	if len(cnt) == 0 {
		return
	}
	return nil
}

```



```
class Solution:
    def findOriginalArray(self, changed: List[int]) -> List[int]:
        length = len(changed)
        changed.sort()
        if length % 2 != 0 :
            return []

        m = {}
        result = []
        for i in changed:
            if i %2 == 0  and i // 2 in m and m[i//2] > 0:
                m[i//2] -= 1
                result.append(i//2)
            else:
                if i not in m:
                    m[i] = 1
                else:
                    m[i] += 1
        for j in m.values():
            if j > 0 :
                return []
        return result
```