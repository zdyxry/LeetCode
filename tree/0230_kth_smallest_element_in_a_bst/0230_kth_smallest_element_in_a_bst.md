230. Kth Smallest Element in a BST

Medium

Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.

Note:

You may assume k is always valid, 1 ≤ k ≤ BST's total elements.

Example 1:

```
Input: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
Output: 1
```

Example 2:

```
Input: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
Output: 3
```

Follow up:

What if the BST is modified (insert/delete operations) often and you need to find the kth smallest frequently? How would you optimize the kthSmallest routine?

## 方法 

二叉搜索树有序的特性，所以中序遍历它，遍历到第 K 个数的时候就是结果

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	res, count := 0, 0
	inorder(root, k, &count, &res)
	return res
}

func inorder(node *TreeNode, k int, count *int, ans *int) {
	if node != nil {
		inorder(node.Left, k, count, ans)
		*count++
		if *count == k {
			*ans = node.Val
			return
		}
		inorder(node.Right, k, count, ans)
	}
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
func kthSmallest(root *TreeNode, k int) int {
    leftSize := getSize(root.Left)
	switch {
	case k <= leftSize:
		// 答案存在于 root.Left 中
		return kthSmallest(root.Left, k)
	case leftSize+1 < k:
		// 答案存在于 root.Right 中
		return kthSmallest(root.Right, k-leftSize-1)
	default:
		// 答案是 root.Val
		return root.Val
	}
}

// 获取 root 树的节点数量
func getSize(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + getSize(root.Left) + getSize(root.Right)
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
    def kthSmallest(self, root, k):
        """
        :type root: TreeNode
        :type k: int
        :rtype: int
        """
        s, cur, rank = [], root, 0

        while s or cur:
            if cur:
                s.append(cur)
                cur = cur.left
            else:
                cur = s.pop()
                rank += 1
                if rank == k:
                    return cur.val
                cur = cur.right

        return float("-inf")
```



