993. Cousins in Binary Tree


Easy


In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.

Two nodes of a binary tree are cousins if they have the same depth, but have different parents.

We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.

Return true if and only if the nodes corresponding to the values x and y are cousins.

 

Example 1:

![1](993-1.png)

Input: root = [1,2,3,4], x = 4, y = 3

Output: false

Example 2:

![2](993-2.png)

Input: root = [1,2,3,null,4,null,5], x = 5, y = 4

Output: true

Example 3:

![3](993-3.png)

Input: root = [1,2,3,null,4], x = 2, y = 3

Output: false
 

Note:

The number of nodes in the tree will be between 2 and 100.

Each node has a unique integer value from 1 to 100.

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
func isCousins(root *TreeNode, x int, y int) bool {
    root = &TreeNode{Left: root}
	px, dx := dfs(root, x)
	py, dy := dfs(root, y)
	return dx == dy && px != py
}

// dfs do NOT check the first root
func dfs(root *TreeNode, x int) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}

	if (root.Left != nil && root.Left.Val == x) ||
		(root.Right != nil && root.Right.Val == x) {
		return root, 1
	}

	if parent, depth := dfs(root.Left, x); depth > 0 {
		return parent, depth + 1
	}

	if parent, depth := dfs(root.Right, x); depth > 0 {
		return parent, depth + 1
	}

	return nil, 0
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
    def isCousins(self, root, x, y):
        """
        :type root: TreeNode
        :type x: int
        :type y: int
        :rtype: bool
        """
        def dfs(root, x, depth, parent):
            if not root:
                return False
            if root.val == x:
                return True
            depth[0] += 1
            prev_parent, parent[0] = parent[0], root
            if dfs(root.left, x, depth, parent):
                return True
            parent[0] = root
            if dfs(root.right, x, depth, parent):
                return True
            parent[0] = prev_parent
            depth[0] -= 1
            return False
        
        depth_x, depth_y = [0], [0]
        parent_x, parent_y = [None], [None]
        return dfs(root, x, depth_x, parent_x) and \
               dfs(root, y, depth_y, parent_y) and \
               depth_x[0] == depth_y[0] and \
               parent_x[0] != parent_y[0]
```