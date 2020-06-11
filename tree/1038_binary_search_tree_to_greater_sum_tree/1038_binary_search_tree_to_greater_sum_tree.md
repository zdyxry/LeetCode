1038. Binary Search Tree to Greater Sum Tree


Medium


Given the root of a binary search tree with distinct values, modify it so that every node has a new value equal to the sum of the values of the original tree that are greater than or equal to node.val.

As a reminder, a binary search tree is a tree that satisfies these constraints:

The left subtree of a node contains only nodes with keys less than the node's key.  
The right subtree of a node contains only nodes with keys greater than the node's key.  
Both the left and right subtrees must also be binary search trees.
 

Example 1:

![1038](1038.png)

```
Input: [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
```

Constraints:

The number of nodes in the tree is between 1 and 100.
Each node will have value between 0 and 100.
The given tree is a binary search tree.


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

func bstToGst(root *TreeNode) *TreeNode {
    sum := 0
    dfs(root, &sum)
    return root
}

func dfs(root *TreeNode, sum *int) {
    if root == nil { return }
    dfs(root.Right, sum)
    *sum += root.Val
    root.Val = *sum
    dfs(root.Left, sum)
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
    def bstToGst(self, root: TreeNode) -> TreeNode:
        if root is None:
            return None
        
        currSum = [0]
        self.recurse(root, currSum)
        return root
        
    def recurse(self, root, currSum):
        if root is None:
            return
        self.recurse(root.right, currSum)
        root.val = root.val + currSum[0]
        currSum[0] = root.val
        self.recurse(root.left, currSum)
```