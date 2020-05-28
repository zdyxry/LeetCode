1457. Pseudo-Palindromic Paths in a Binary Tree


Medium


Given a binary tree where node values are digits from 1 to 9. A path in the binary tree is said to be pseudo-palindromic if at least one permutation of the node values in the path is a palindrome.

Return the number of pseudo-palindromic paths going from the root node to leaf nodes.

 

Example 1:



```
Input: root = [2,3,1,3,1,null,1]
Output: 2 
Explanation: The figure above represents the given binary tree. There are three paths going from the root node to leaf nodes: the red path [2,3,3], the green path [2,1,1], and the path [2,3,1]. Among these paths only red path and green path are pseudo-palindromic paths since the red path [2,3,3] can be rearranged in [3,2,3] (palindrome) and the green path [2,1,1] can be rearranged in [1,2,1] (palindrome).
```

Example 2:



```
Input: root = [2,1,1,1,3,null,null,null,null,null,1]
Output: 1 
Explanation: The figure above represents the given binary tree. There are three paths going from the root node to leaf nodes: the green path [2,1,1], the path [2,1,3,1], and the path [2,1]. Among these paths only the green path is pseudo-palindromic since [2,1,1] can be rearranged in [1,2,1] (palindrome).
```

Example 3:

```
Input: root = [9]
Output: 1
```
 

Constraints:

The given binary tree will have between 1 and 10^5 nodes.  
Node values are digits from 1 to 9.


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
func pseudoPalindromicPaths (root *TreeNode) int {
    if root == nil {
        return 0
    }
    var ans int
    var h func(node *TreeNode, temp int)
    h = func(node *TreeNode, temp int) {
        temp ^= (1<<node.Val)
        if node.Left == nil && node.Right == nil {
            if temp == 0 || temp & (temp-1) == 0 {
                ans++
            }
        }

        if node.Left != nil {
            h(node.Left, temp)
        }

        if node.Right != nil {
            h(node.Right, temp)
        }

    }

    h(root, 0)
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
    def pseudoPalindromicPaths (self, root: TreeNode) -> int:
        s = set()

        def dfs(root):

            if not root: return 0

            if root.val in s:
                s.discard(root.val)
            else:
                s.add(root.val)

            res = dfs(root.left) + dfs(root.right)

            if not root.left and not root.right:
                res += len(s)<=1

            if root.val in s:
                s.discard(root.val)
            else:
                s.add(root.val)

            return res

        return dfs(root)
```