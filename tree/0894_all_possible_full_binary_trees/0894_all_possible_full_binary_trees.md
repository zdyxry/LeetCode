894. All Possible Full Binary Trees


Medium


A full binary tree is a binary tree where each node has exactly 0 or 2 children.

Return a list of all possible full binary trees with N nodes.  Each element of the answer is the root node of one possible tree.

Each node of each tree in the answer must have node.val = 0.

You may return the final list of trees in any order.

 

Example 1:

```
Input: 7
Output: [[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
Explanation:
```
 

Note:

1 <= N <= 20


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
func allPossibleFBT(N int) []*TreeNode {
	if N%2 == 0 {
		return nil
	}
	if N == 1 {
		return []*TreeNode{new(TreeNode)}
	}

	trees := make([]*TreeNode, 0)
	for i := 1; i <= N-2; i=i+2 {   //注意循环的起点和循环条件
		l := allPossibleFBT(i)
		r := allPossibleFBT(N-1-i)
		for _, left := range l {
			for _, right := range r {
				node := new(TreeNode)
				node.Left=left
				node.Right=right
				trees = append(trees, node)
			}
		}
	}
	return trees
}

```


```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    def allPossibleFBT(self, N: int) -> List[TreeNode]:
        res = []
        if N == 1:
            return [TreeNode(0)]
        if N % 2 == 0:
            return []
        
        left_num = 1
        right_num = N - 2
        
        while right_num > 0:
            left_tree = self.allPossibleFBT(left_num)
            right_tree = self.allPossibleFBT(right_num)
            for i in range(len(left_tree)):
                for j in range(len(right_tree)):
                    root = TreeNode(0)
                    root.left = left_tree[i]
                    root.right = right_tree[j]
                    res.append(root)
            left_num += 2
            right_num -= 2
        
        return res     
```