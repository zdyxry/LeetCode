145. Binary Tree Postorder Traversal


Medium


Given the root of a binary tree, return the postorder traversal of its nodes' values.

 

Example 1:


```
Input: root = [1,null,2,3]
Output: [3,2,1]
```

Example 2:

```
Input: root = []
Output: []
```

Example 3:

```
Input: root = [1]
Output: [1]
```

Example 4:


```
Input: root = [1,2]
Output: [2,1]
```

Example 5:


```
Input: root = [1,null,2]
Output: [2,1]
```
 

Constraints:

The number of the nodes in the tree is in the range [0, 100].    
-100 <= Node.val <= 100
 

Follow up:

Recursive solution is trivial, could you do it iteratively?



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
func postorderTraversal(root *TreeNode) (res []int) {
    var postorder func(*TreeNode)
    postorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        postorder(node.Left)
        postorder(node.Right)
        res = append(res, node.Val)
    }
    postorder(root)
    return
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
    def postorderTraversal(self, root: TreeNode) -> List[int]:
        def postorder(root: TreeNode):
            if not root:
                return
            postorder(root.left)
            postorder(root.right)
            res.append(root.val)
        
        res = list()
        postorder(root)
        return res
```