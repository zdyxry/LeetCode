114. Flatten Binary Tree to Linked List


Medium


Given a binary tree, flatten it to a linked list in-place.

For example, given the following tree:

```
    1
   / \
  2   5
 / \   \
3   4   6
The flattened tree should look like:

1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
```

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
func flatten(root *TreeNode)  {
    recur(root)
}

// recur 会 flatten root，并返回 flatten 后的 leaf 节点
func recur(root *TreeNode) *TreeNode {
	if root == nil ||
		(root.Left == nil && root.Right == nil) {
		return root
	}

	if root.Left == nil {
		return recur(root.Right)
	}

	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		return recur(root.Right)
	}

	res := recur(root.Right)
	recur(root.Left).Right = root.Right
	root.Right = root.Left
	// 不要忘记封闭 left 节点
	root.Left = nil

	return res
}
```



```python
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def flatten(self, root):
        """
        :type root: TreeNode
        :rtype: None Do not return anything, modify root in-place instead.
        """
        self.flattenRecu(root, None)

    def flattenRecu(self, root, list_head):
        if root:
            list_head = self.flattenRecu(root.right, list_head)
            list_head = self.flattenRecu(root.left, list_head)
            root.right = list_head
            root.left = None
            return root
        else:
            return list_head
```