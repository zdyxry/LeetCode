653. Two Sum IV - Input is a BST


Easy


Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

Example 1:

```
Input: 
    5
   / \
  3   6
 / \   \
2   4   7

Target = 9

Output: True
```

Example 2:

```
Input: 
    5
   / \
  3   6
 / \   \
2   4   7

Target = 28

Output: False
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

func findTarget(root *TreeNode, k int) bool {
	return helper(root, root, k)
}

func helper(root, searchRoot *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	return (root.Val*2 != k && findNode(searchRoot, k-root.Val)) ||
		helper(root.Left, searchRoot, k) ||
		helper(root.Right, searchRoot, k)
}

func findNode(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	if root.Val == target {
		return true
	}

	if root.Val < target {
		return findNode(root.Right, target)
	}
	return findNode(root.Left, target)
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
    def findTarget(self, root, k):
        """
        :type root: TreeNode
        :type k: int
        :rtype: bool
        """
        res = set()
        def inOrder(root):
            if not root:
                return False
            if k - root.val in res:
                return True
            res.add(root.val)
            return inOrder(root.left) or inOrder(root.right)
        return inOrder(root)

```

```python 
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution(object):
    def findTarget(self, root, k):
        """
        :type root: TreeNode
        :type k: int
        :rtype: bool
        """
        if not root: return False
        bfs, s = [root], set()
        for i in bfs:
            if k - i.val in s : return True
            s.add(i.val)
            if i.left : bfs.append(i.left)
            if i.right : bfs.append(i.right)
        return False

```