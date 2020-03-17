1382. Balance a Binary Search Tree


Medium


Given a binary search tree, return a balanced binary search tree with the same node values.

A binary search tree is balanced if and only if the depth of the two subtrees of every node never differ by more than 1.

If there is more than one answer, return any of them.

 

Example 1:

![1382](1382.png)

```
Input: root = [1,null,2,null,3,null,4,null,null]
Output: [2,1,3,null,null,null,4]
Explanation: This is not the only correct answer, [3,1,4,null,2,null,null] is also correct.
```

Constraints:

The number of nodes in the tree is between 1 and 10^4.  
The tree nodes will have distinct values between 1 and 10^5.

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
func balanceBST(root *TreeNode) *TreeNode {
    sortedArray := getSortedArray(root)
	return buildBalanceBySortedArray(sortedArray)

}
func buildBalanceBySortedArray(array []int) *TreeNode {
	if len(array) == 0 {
		return nil
	}

	mid := len(array) / 2
	root := &TreeNode{Val: array[mid]}
	root.Left = buildBalanceBySortedArray(array[:mid])
	root.Right = buildBalanceBySortedArray(array[mid+1:])
	return root
}

func getSortedArray(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	res = append(res, getSortedArray(root.Left)...)
	res = append(res, root.Val)
	res = append(res, getSortedArray(root.Right)...)
	return res
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
    def balanceBST(self, root):
        """
        :type root: TreeNode
        :rtype: TreeNode
        """
        def dfs(node):
            """inorder depth-first traverse bst"""
            if not node: return 
            dfs(node.left)
            value.append(node.val)
            dfs(node.right)
        
        value = [] #collect values
        dfs(root)
        if not value: return None
        
        def tree(x): 
            if not x: return None
            k = len(x)//2
            ans = TreeNode(x[k])
            ans.left = tree(x[:k])
            ans.right = tree(x[k+1:])
            return ans
        
        return tree(value)
```