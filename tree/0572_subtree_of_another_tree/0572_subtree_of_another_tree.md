572. Subtree of Another Tree


Easy


Given two non-empty binary trees s and t, check whether tree t has exactly the same structure and node values with a subtree of s. A subtree of s is a tree consists of a node in s and all of this node's descendants. The tree s could also be considered as a subtree of itself.

Example 1:

```
Given tree s:

     3
    / \
   4   5
  / \
 1   2
```

```
Given tree t:
   4 
  / \
 1   2
Return true, because t has the same structure and node values with a subtree of s.
```


Example 2:

```
Given tree s:

     3
    / \
   4   5
  / \
 1   2
    /
   0
```

```
Given tree t:
   4
  / \
 1   2
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
func isSubtree(s *TreeNode, t *TreeNode) bool {
    if isSame(s, t) { return true }
    
    if s == nil {return false}
    return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSame(s *TreeNode, t *TreeNode) bool {
    if s == nil {return t == nil}
    if t == nil {return false}
    
    if s.Val != t.Val {
        return false
    }
    
    return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
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
    def isSubtree(self, s, t):
        """
        :type s: TreeNode
        :type t: TreeNode
        :rtype: bool
        """
        def isSame(x, y):
            if not x and not y:
                return True
            if not x or not y:
                return False
            return x.val == y.val and \
                   isSame(x.left, y.left) and \
                   isSame(x.right, y.right)

        def preOrderTraverse(s, t):
            return s != None and \
                   (isSame(s, t) or \
                    preOrderTraverse(s.left, t) or \
                    preOrderTraverse(s.right, t))

        return preOrderTraverse(s, t)
```