1530. Number of Good Leaf Nodes Pairs


Medium


Given the root of a binary tree and an integer distance. A pair of two different leaf nodes of a binary tree is said to be good if the length of the shortest path between them is less than or equal to distance.

Return the number of good leaf node pairs in the tree.

 

Example 1:


```
Input: root = [1,2,3,null,4], distance = 3
Output: 1
Explanation: The leaf nodes of the tree are 3 and 4 and the length of the shortest path between them is 3. This is the only good pair.
```

Example 2:


```
Input: root = [1,2,3,4,5,6,7], distance = 3
Output: 2
Explanation: The good pairs are [4,5] and [6,7] with shortest path = 2. The pair [4,6] is not good because the length of ther shortest path between them is 4.
```

Example 3:

```
Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2], distance = 3
Output: 1
Explanation: The only good pair is [2,5].
```

Example 4:

```
Input: root = [100], distance = 1
Output: 0
```

Example 5:

```
Input: root = [1,1,1], distance = 2
Output: 1
```
 

Constraints:

The number of nodes in the tree is in the range [1, 2^10].  
Each node's value is between [1, 100].  
1 <= distance <= 10


## 方法

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func dfs(n *TreeNode, d int) (map[int]int, int) {
    if n == nil {
        return map[int]int{}, 0
    }
    if n.Left == nil && n.Right == nil {
        return map[int]int{0: 1}, 0
    }

    lp, lc := dfs(n.Left, d)
    rp, rc := dfs(n.Right, d)
    
    m := map[int]int{}
    c := lc + rc 
    for kr, vr := range rp {
        m[kr + 1] += vr
    }
    
    for kl, vl := range lp {
        m[kl + 1] += vl
        for kr, vr := range rp {
            if kl + kr + 2 <= d {
                c += vl * vr
            }    
        }    
    }
    
    //fmt.Println(n.Val, m, c)
    
    return m, c
}

func countPairs(root *TreeNode, distance int) int {
    _, lc := dfs(root, distance)
    return lc
}

```

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def countPairs(self, root: TreeNode, distance: int) -> int:
        if not root: return 0
        self.res = 0
        def dfs(node, depth):
            if not node.left and not node.right: return collections.Counter([depth])
            l, r = collections.Counter(), collections.Counter()
            if node.left:
                l = dfs(node.left, depth + 1)
            if node.right:
                r = dfs(node.right, depth + 1)
            for k in l:
                for j in r:
                    if j + k <= distance +2 * depth:
                        self.res += l[k] * r[j]
            return l + r
        dfs(root, 0)
        return self.res
```