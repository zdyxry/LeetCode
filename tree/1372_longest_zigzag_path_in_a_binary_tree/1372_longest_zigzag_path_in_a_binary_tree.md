1372. Longest ZigZag Path in a Binary Tree


Medium


Given a binary tree root, a ZigZag path for a binary tree is defined as follow:

Choose any node in the binary tree and a direction (right or left).

If the current direction is right then move to the right child of the current node otherwise move to the left child.

Change the direction from right to left or right to left.

Repeat the second and third step until you can't move in the tree.

Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).

Return the longest ZigZag path contained in that tree.

 

Example 1:

![1](1372-1.png)

```
Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]
Output: 3
Explanation: Longest ZigZag path in blue nodes (right -> left -> right).
```

Example 2:

![2](1372-2.png)


```
Input: root = [1,1,1,null,1,null,null,1,1,null,1]
Output: 4
Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).
```

Example 3:

```
Input: root = [1]
Output: 0
```

Constraints:

Each tree has at most 50000 nodes..

Each node's value is between [1, 100].

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
var res int
func longestZigZag(root *TreeNode) int {
    res = 0
    if root == nil {
        return 0
    }
    traverse(root)
    return res
}

func traverse(root *TreeNode) (int, int){
    l, r := 1, 1
    if root.Left != nil {
        _, lr := traverse(root.Left)
        l += lr
        res = max(res, l - 1)
    }
    if root.Right != nil {
        rl, _ := traverse(root.Right)
        r += rl
        res = max(res, r - 1)
    }
    return l, r
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
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
func longestZigZag(root *TreeNode) int {
    if root == nil {
        return 0
    }
    res, _, _ := bst(root)
    return res - 1
}

func bst(node *TreeNode) (int, int, int) {
    if node == nil {
        return 0, 0, 0
    }
    lb, _, lr := bst(node.Left)
    rb, rl, _ := bst(node.Right)
    l := lr + 1
    r := rl + 1
    b := max(max(l, r), max(lb, rb))
    return b, l, r
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
    def longestZigZag(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        def dfs(root):
            if not root:
                return [-1, -1, -1]
            left, right = dfs(root.left), dfs(root.right)
            return [left[1] + 1, right[0] + 1, max(left[1]+1, right[0]+1, left[2], right[2])]
        return dfs(root)[-1]
```