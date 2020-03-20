897. Increasing Order Search Tree


Easy


Given a binary search tree, rearrange the tree in in-order so that the leftmost node in the tree is now the root of the tree, and every node has no left child and only 1 right child.

Example 1:

```
Input: [5,3,6,2,4,null,8,1,null,null,null,7,9]

       5
      / \
    3    6
   / \    \
  2   4    8
 /        / \ 
1        7   9

Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

 1
  \
   2
    \
     3
      \
       4
        \
         5
          \
           6
            \
             7
              \
               8
                \
                 9  
```

Constraints:

The number of nodes in the given tree will be between 1 and 100.   
Each node will have a unique integer value from 0 to 1000.


## 方法


```go
func increasingBST(root *TreeNode) *TreeNode {
	vals := []int{}
	increasingBSTRecursive(root, &vals)
	ans := &TreeNode{Val:0}
	cur := ans
	for _, v := range vals {
		cur.Right = &TreeNode{Val:v}
		cur = cur.Right
	}
	return ans.Right
}

func increasingBSTRecursive(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	increasingBSTRecursive(root.Left, vals)
	*vals = append(*vals, root.Val)
	increasingBSTRecursive(root.Right, vals)
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
func increasingBST(root *TreeNode) *TreeNode {
    min, _ := helper(root)
    return min
}

func helper(node *TreeNode) (*TreeNode, *TreeNode) {
    min, max := node, node
    if node.Left != nil {
        lMin, lMax := helper(node.Left)
        min = lMin
        lMax.Right = node
    }
    if node.Right != nil {
        rMin, rMax := helper(node.Right)
        max = rMax
        node.Right = rMin
    }
    node.Left = nil
    return min, max
}
```

```python
class Solution:
    def increasingBST(self, root):
        def inorder(node):
            if node:
                inorder(node.left)
                node.left = None
                self.cur.right = node
                self.cur = node
                inorder(node.right)

        ans = self.cur = TreeNode(None)
        inorder(root)
        return ans.right
```



```python
class Solution:
    def increasingBST(self, root):
        def inorder(node):
            if node:
                yield from inorder(node.left)
                yield node.val
                yield from inorder(node.right)

        ans = cur = TreeNode(None)
        for v in inorder(root):
            cur.right = TreeNode(v)
            cur = cur.right
        return ans.right
```