110. Balanced Binary Tree


Easy


Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example 1:

```
Given the following tree [3,9,20,null,null,15,7]:

    3
   / \
  9  20
    /  \
   15   7
Return true.
```

Example 2:

```
Given the following tree [1,2,2,3,3,null,null,4,4]:

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
Return false.
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

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHight := depth(root.Left)
	rightHight := depth(root.Right)
	return abs(leftHight-rightHight) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func max(a,b int) int {
    if a > b {
        return a
    } else {
        return b
    }
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
    def isBalanced(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        def getDepth(node):
            if not node: return 0
            left = getDepth(node.left)
            right = getDepth(node.right)
            if abs(left - right) > 1: self.res = False
            return 1 + max(left, right)

        self.res = True
        getDepth(root)
        return self.res
```