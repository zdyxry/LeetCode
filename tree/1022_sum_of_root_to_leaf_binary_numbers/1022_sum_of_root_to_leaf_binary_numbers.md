1022. Sum of Root To Leaf Binary Numbers


Easy


Given a binary tree, each node has value 0 or 1.  Each root-to-leaf path represents a binary number starting with the most significant bit.  For example, if the path is 0 -> 1 -> 1 -> 0 -> 1, then this could represent 01101 in binary, which is 13.

For all leaves in the tree, consider the numbers represented by the path from the root to that leaf.

Return the sum of these numbers.

 

Example 1:



```
Input: [1,0,1,0,1,0,1]
Output: 22
Explanation: (100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
```

Note:

The number of nodes in the tree is between 1 and 1000.  
node.val is 0 or 1.  
The answer will not exceed 2^31 - 1.  


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
func sumRootToLeaf(root *TreeNode) int {
    ans := 0
    var f func(*TreeNode, int); f = func(r *TreeNode, s int) {
        if r != nil {
            s = s * 2 + r.Val
            if r.Left == nil && r.Right == nil {
                ans += s
            } else {
                f(r.Left, s)
                f(r.Right, s)
            }
        }
    }
    f(root, 0)
    return ans
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
    def sumRootToLeaf(self, root: TreeNode) -> int:
        return self.helper(root, 0)
    def helper(self, root, base):
        if root == None:
            return 0
        if root.left == None and root.right == None:
            return 2 * base + root.val
        base = 2 * base + root.val
        return self.helper(root.left, base) + self.helper(root.right, base)
```