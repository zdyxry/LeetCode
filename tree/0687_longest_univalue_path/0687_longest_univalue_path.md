687. Longest Univalue Path


Easy


Given a binary tree, find the length of the longest path where each node in the path has the same value. This path may or may not pass through the root.

The length of path between two nodes is represented by the number of edges between them.

 

Example 1:

```
Input:

              5
             / \
            4   5
           / \   \
          1   1   5
Output: 2

```
 

Example 2:

```
Input:

              1
             / \
            4   5
           / \   \
          4   4   5
Output: 2
```
 

Note: The given binary tree has not more than 10000 nodes. The height of the tree is not more than 1000.


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
func longestUnivaluePath(root *TreeNode) int {
    maxLen := 0
	helper(root, &maxLen)
	return maxLen
}

// 返回从 root 出发拥有相同 Val 值的线路上的 edge 数量
func helper(n *TreeNode, maxLen *int) int {
	if n == nil {
		return 0
	}

	l := helper(n.Left, maxLen)
	r := helper(n.Right, maxLen)
	res := 0

	// 左侧单边的最长距离
	if n.Left != nil && n.Val == n.Left.Val {
		*maxLen = max(*maxLen, l+1)
		res = max(res, l+1)
	}
	// 右侧单边的最长距离
	if n.Right != nil && n.Val == n.Right.Val {
		*maxLen = max(*maxLen, r+1)
		res = max(res, r+1)
	}
	// 通过根节点的最长边
	if n.Left != nil && n.Val == n.Left.Val &&
		n.Right != nil && n.Val == n.Right.Val {
		*maxLen = max(*maxLen, l+r+2)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
    def longestUnivaluePath(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        result = [0]
        def dfs(node):
            if not node:
                return 0
            left, right = dfs(node.left), dfs(node.right)
            left = (left+1) if node.left and node.left.val == node.val else 0
            right = (right+1) if node.right and node.right.val == node.val else 0
            result[0] = max(result[0], left+right)
            return max(left, right)

        dfs(root)
        return result[0]
```