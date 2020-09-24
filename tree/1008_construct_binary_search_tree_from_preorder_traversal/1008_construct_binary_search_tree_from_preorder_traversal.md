1008. Construct Binary Search Tree from Preorder Traversal


Medium


Return the root node of a binary search tree that matches the given preorder traversal.

(Recall that a binary search tree is a binary tree where for every node, any descendant of node.left has a value < node.val, and any descendant of node.right has a value > node.val.  Also recall that a preorder traversal displays the value of the node first, then traverses node.left, then traverses node.right.)

It's guaranteed that for the given test cases there is always possible to find a binary search tree with the given requirements.

Example 1:

```
Input: [8,5,1,7,10,12]
Output: [8,5,10,1,7,null,12]
```
 

Constraints:

1 <= preorder.length <= 100  
1 <= preorder[i] <= 10^8  
The values of preorder are distinct.


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
func bstFromPreorder(preorder []int) *TreeNode {
    if devide := len(preorder); devide > 0 {
        for i, val := range preorder {
            if val > preorder[0] {
                devide = i
                break
            }
        }
        return &TreeNode{
            Val: preorder[0],
            Left: bstFromPreorder(preorder[1: devide]),
            Right: bstFromPreorder(preorder[devide: ]),
        }
    }
    return nil
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
    def bstFromPreorder(self, preorder: List[int]) -> TreeNode:
        if len(preorder) == 0:
            return None
        leftList = []
        rightList = []
        root = TreeNode(preorder[0])
        for num in preorder[1:]:
            if num < preorder[0]:
                leftList.append(num)
            elif num > preorder[0]:
                rightList.append(num)
        root.left = self.bstFromPreorder(leftList)
        root.right = self.bstFromPreorder(rightList)
        return root
```