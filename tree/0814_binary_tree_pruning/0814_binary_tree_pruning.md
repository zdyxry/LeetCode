814. Binary Tree Pruning


Medium


We are given the head node root of a binary tree, where additionally every node's value is either a 0 or a 1.

Return the same tree where every subtree (of the given tree) not containing a 1 has been removed.

(Recall that the subtree of a node X is X, plus every node that is a descendant of X.)

Example 1:
```
Input: [1,null,0,0,1]
Output: [1,null,0,null,1]

 
Explanation: 
Only the red nodes satisfy the property "every subtree not containing a 1".
The diagram on the right represents the answer.
```

Example 2:
```
Input: [1,0,1,0,0,0,1]
Output: [1,null,1,null,1]
```


Example 3:
```
Input: [1,1,0,1,1,0,1,0]
Output: [1,1,0,1,1,null,1]
```


Note:

The binary tree will have at most 200 nodes.  
The value of each node will only be 0 or 1.


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
func pruneTree(root *TreeNode) *TreeNode {
	return deal(root)
}

func deal(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	node.Left = deal(node.Left)
	node.Right = deal(node.Right)
	if node.Left == nil && node.Right == nil && node.Val == 0 {
		return nil
	}
	return node
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
    def pruneTree(self, root: TreeNode) -> TreeNode:
        if root is None:
            return root
        root.left = self.pruneTree(root.left)
        root.right = self.pruneTree(root.right)
        if root.left is None and root.right is None and root.val == 0:
            return None
        return root
```