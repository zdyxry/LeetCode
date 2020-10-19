1110. Delete Nodes And Return Forest


Medium


Given the root of a binary tree, each node in the tree has a distinct value.

After deleting all nodes with a value in to_delete, we are left with a forest (a disjoint union of trees).

Return the roots of the trees in the remaining forest.  You may return the result in any order.

 

Example 1:

![1](1110-1.png)

```
Input: root = [1,2,3,4,5,6,7], to_delete = [3,5]
Output: [[1,2,null,4],[6],[7]]
```

Constraints:

```
The number of nodes in the given tree is at most 1000.
Each node has a distinct value between 1 and 1000.
to_delete.length <= 1000
to_delete contains distinct values between 1 and 1000.
```

##  方法



```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
    m := make(map[int]bool)
    for i := 0; i < len(to_delete); i++ {
        m[to_delete[i]] = true
    }
    return delNode(root, m)
}

func delNode(root *TreeNode, m map[int]bool) []*TreeNode {
    root, ret := del(root, m)
    if root == nil {
        return ret
    }
    return append(ret, root)
}

func del(root *TreeNode, m map[int]bool) (*TreeNode, []*TreeNode) {
    ret := []*TreeNode{}
    if root == nil {
        return nil, nil
    }
    if m[root.Val] {
        return nil, append(delNode(root.Left, m), delNode(root.Right, m)...)
    }
    left, l := del(root.Left, m)
    root.Left = left
    ret = append(ret, l...)
    
    right, r := del(root.Right, m)
    root.Right = right
    ret = append(ret, r...)
    
    return root, ret
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
    def delNodes(self, root: TreeNode, to_delete: List[int]) -> List[TreeNode]:
        if not root: return None
        mapper = set(to_delete)
        ans = [] if root.val in mapper else [root]
        def dfs(node, parent, direction):
            if not node: return
            dfs(node.left, node, 'left')
            dfs(node.right, node, 'right')
            if node.val in mapper:
                if node.left: ans.append(node.left)
                if node.right: ans.append(node.right)
                if direction == 'left':
                    parent.left = None
                if direction == 'right':
                    parent.right = None
        dfs(root, None, None)
        return ans
```