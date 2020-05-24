1452. People Whose List of Favorite Companies Is Not a Subset of Another List


Medium


Given the array favoriteCompanies where favoriteCompanies[i] is the list of favorites companies for the ith person (indexed from 0).

Return the indices of people whose list of favorite companies is not a subset of any other list of favorites companies. You must return the indices in increasing order.

 

Example 1:

```
Input: favoriteCompanies = [["leetcode","google","facebook"],["google","microsoft"],["google","facebook"],["google"],["amazon"]]
Output: [0,1,4] 
Explanation: 
Person with index=2 has favoriteCompanies[2]=["google","facebook"] which is a subset of favoriteCompanies[0]=["leetcode","google","facebook"] corresponding to the person with index 0. 
Person with index=3 has favoriteCompanies[3]=["google"] which is a subset of favoriteCompanies[0]=["leetcode","google","facebook"] and favoriteCompanies[1]=["google","microsoft"]. 
Other lists of favorite companies are not a subset of another list, therefore, the answer is [0,1,4].
```

Example 2:

```
Input: favoriteCompanies = [["leetcode","google","facebook"],["leetcode","amazon"],["facebook","google"]]
Output: [0,1] 
Explanation: In this case favoriteCompanies[2]=["facebook","google"] is a subset of favoriteCompanies[0]=["leetcode","google","facebook"], therefore, the answer is [0,1].
```

Example 3:

```
Input: favoriteCompanies = [["leetcode"],["google"],["facebook"],["amazon"]]
Output: [0,1,2,3]
```
 

Constraints:

1 <= favoriteCompanies.length <= 100  
1 <= favoriteCompanies[i].length <= 500  
1 <= favoriteCompanies[i][j].length <= 20  
All strings in favoriteCompanies[i] are distinct.  
All lists of favorite companies are distinct, that is, If we sort alphabetically each list then favoriteCompanies[i] != favoriteCompanies[j].  
All strings consist of lowercase English letters only.

## 方法

```go
func peopleIndexes(favoriteCompanies [][]string) []int {
    ans := make([]int, 0)
	dist := make(map[int][]map[string]int)
	for _, v := range favoriteCompanies {
		// v - []string
		m := make(map[string]int)
		for _, val := range v {
			// val - company name str
			m[val]++
		}
		dist[len(v)] = append(dist[len(v)], m)
	}
	for k, v := range favoriteCompanies {
		if subset(v, dist) {
			ans = append(ans, k)
		}
	}
	return ans
}

func subset(d1 []string, d2 map[int][]map[string]int) bool {
	ans := 0
	for k, v := range d2 {
		if k >= len(d1) {
			for _, m := range v {
				sub := true
				for _, s := range d1 {
					if _, exist := m[s]; !exist {
						sub = false
						break
					}
				}
				if sub {
					ans++
				}
			}
		}
	}
	if ans >= 2 {
		return false
	} else {
		return true
	}
}
```


```python
class Solution:
    def peopleIndexes(self, favoriteCompanies: List[List[str]]) -> List[int]:
        l=favoriteCompanies
        for i in range(len(l)):
            l[i]=[l[i],i]
        l.sort(key=lambda x:len(x[0]))
        ans=[]
        for i in range(len(l)):
            for j in range(i+1,len(l)):
                if set(l[i][0]).issubset(set(l[j][0])):
                    break
            else:
                ans.append(l[i][1])
        ans.sort()
        return ans
```