1282. Group the People Given the Group Size They Belong To


Medium


There are n people that are split into some unknown number of groups. Each person is labeled with a unique ID from 0 to n - 1.

You are given an integer array groupSizes, where groupSizes[i] is the size of the group that person i is in. For example, if groupSizes[1] = 3, then person 1 must be in a group of size 3.

Return a list of groups such that each person i is in a group of size groupSizes[i].

Each person should appear in exactly one group, and every person must be in a group. If there are multiple answers, return any of them. It is guaranteed that there will be at least one valid solution for the given input.

 

Example 1:

```
Input: groupSizes = [3,3,3,3,3,1,3]
Output: [[5],[0,1,2],[3,4,6]]
Explanation: 
The first group is [5]. The size is 1, and groupSizes[5] = 1.
The second group is [0,1,2]. The size is 3, and groupSizes[0] = groupSizes[1] = groupSizes[2] = 3.
The third group is [3,4,6]. The size is 3, and groupSizes[3] = groupSizes[4] = groupSizes[6] = 3.
Other possible solutions are [[2,1,6],[5],[0,4,3]] and [[5],[0,6,2],[4,3,1]].
```

Example 2:

```
Input: groupSizes = [2,1,3,3,3,2]
Output: [[1],[0,5],[2,3,4]]
```
 

Constraints:

groupSizes.length == n   
1 <= n <= 500   
1 <= groupSizes[i] <= n


## 方法


```go
func groupThePeople(groupSizes []int) [][]int {
    var groups [][]int
	groupMemberCount := make(map[int][]int)

	for i := 0; i < len(groupSizes); i++ {
		groupMemberCount[groupSizes[i]] = append(groupMemberCount[groupSizes[i]], i)
	}

	for groupSize, groupMembers := range groupMemberCount {
		subGroup := []int{}
		for i := 0; i < len(groupMembers); i++ {
			if i%groupSize == groupSize-1 {
				subGroup = append(subGroup, groupMembers[i])
				groups = append(groups, subGroup)
				subGroup = []int{}
				continue
			}
			subGroup = append(subGroup, groupMembers[i])
		}
	}
	return groups
}
```


```python
class Solution:
    def groupThePeople(self, groupSizes: List[int]) -> List[List[int]]:
        groups = collections.defaultdict(list)
        for i, _id in enumerate(groupSizes):
            groups[_id].append(i)
        
        ans = list()
        for gsize, users in groups.items():
            for it in range(0, len(users), gsize):
                ans.append(users[it : it + gsize])
        return ans
```