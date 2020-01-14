102. Binary Tree Level Order Traversal


Medium


Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:

```
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
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
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		// 出现了新的 level
		if  level >= len(res)   {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}
```


```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	curNum, nextLevelNum, res, tmp := 1, 0, [][]int{}, []int{}
	for len(queue) != 0 {
		if curNum > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevelNum++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevelNum++
			}
			curNum--
			tmp = append(tmp, node.Val)
			queue = queue[1:]
		}
		if curNum == 0 {
			res = append(res, tmp)
			curNum = nextLevelNum
			nextLevelNum = 0
			tmp = []int{}
		}
	}
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
    def levelOrder(self, root):
        """
        :type root: TreeNode
        :rtype: List[List[int]]
        """
        if root is None:
            return []
        result, current = [], [root]
        while current:
            next_level, vals = [], []
            for node in current:
                vals.append(node.val)
                if node.left:
                    next_level.append(node.left)
                if node.right:
                    next_level.append(node.right)
            current = next_level
            result.append(vals)
        return result
```