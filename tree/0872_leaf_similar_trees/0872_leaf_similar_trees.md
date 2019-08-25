872. Leaf-Similar Trees


Easy


Consider all the leaves of a binary tree.  From left to right order, the values of those leaves form a leaf value sequence.

![872](./872.png)

For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

Two binary trees are considered leaf-similar if their leaf value sequence is the same.


Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

 

Note:

Both of the given trees will have between 1 and 100 nodes.


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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    a1 := [100]int{}
	s1 := a1[:0]
	search(root1, &s1)

	a2 := [100]int{}
	s2 := a2[:0]
	search(root2, &s2)

	return a1 == a2
}

func search(root *TreeNode, sp *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*sp = append(*sp, root.Val)
		return
	}
	search(root.Left, sp)
	search(root.Right, sp)
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
    def leafSimilar(self, root1, root2):
        """
        :type root1: TreeNode
        :type root2: TreeNode
        :rtype: bool
        """
        return self.iterative(root1,[]) == self.iterative(root2,[])
        
    def iterative(self,root,s):
        if root is not None:
            stack = []
            stack.append(root)
            while stack:
                x = stack.pop(-1)
                if x.left is None and x.right is None:
                    s.append(x.val)
                    continue
                if x.right is not None:
                    stack.append(x.right)
                if x.left is not None:
                    stack.append(x.left)
        return s
```