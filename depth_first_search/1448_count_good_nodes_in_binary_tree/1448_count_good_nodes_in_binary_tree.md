1448. Count Good Nodes in Binary Tree


Medium


Given a binary tree root, a node X in the tree is named good if in the path from root to X there are no nodes with a value greater than X.

Return the number of good nodes in the binary tree.

 

Example 1:

![1](1448-1.png)

```
Input: root = [3,1,4,3,null,1,5]
Output: 4
Explanation: Nodes in blue are good.
Root Node (3) is always a good node.
Node 4 -> (3,4) is the maximum value in the path starting from the root.
Node 5 -> (3,4,5) is the maximum value in the path
Node 3 -> (3,1,3) is the maximum value in the path.
```

Example 2:

![2](1448-2.png)

```
Input: root = [3,3,null,4,2]
Output: 3
Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.
```

Example 3:

```
Input: root = [1]
Output: 1
Explanation: Root is considered as good.
```
 

Constraints:

The number of nodes in the binary tree is in the range [1, 10^5].  
Each node's value is between [-10^4, 10^4].


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
func goodNodes(root *TreeNode) (goodNodeCount int) {
    getGoodNodes(root, root.Val, &goodNodeCount)
    return
}

func getGoodNodes(root *TreeNode, prev int, count *int) {
    if root == nil {
        return
    }
    
    if prev <= root.Val {
        *count+=1
        prev = root.Val
    }
    
    getGoodNodes(root.Left, prev, count)
    getGoodNodes(root.Right, prev, count)   
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
    def goodNodes(self, root: TreeNode) -> int:
        self.res = 0

        def dfs(node, curmx):
            if not node:
                return
            if node.val >= curmx:
                self.res += 1
            curmx = max(curmx, node.val)
            dfs(node.left, curmx)
            dfs(node.right, curmx)

        dfs(root, -float('inf'))
        return self.res
```