437. Path Sum III


Easy


You are given a binary tree in which each node contains an integer value.

Find the number of paths that sum to a given value.

The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).

The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

Example:

```
root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8

      10
     /  \
    5   -3
   / \    \
  3   2   11
 / \   \
3  -2   1

Return 3. The paths that sum to 8 are:

1.  5 -> 3
2.  5 -> 2 -> 1
3. -3 -> 11
```


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
func pathSum(root *TreeNode, sum int) int {
    m := make(map[int]int)
	m[0] = 1
	ans := 0
	helper(root, sum, m, 0, &ans)
	return ans
}

func helper(root *TreeNode, target int, m map[int]int, cur int, ans *int) {
	if root == nil {
		return
	}
	cur = cur + root.Val
	*ans += m[cur-target]
	m[cur]++
    if root.Left != nil {
        helper(root.Left, target, m, cur, ans)    
    }
	
    if root.Right != nil {
        helper(root.Right, target, m, cur, ans)    
    }
	
	m[cur]--
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
func pathSum(root *TreeNode, sum int) int {
    if root == nil {
		return 0
	}

	cnt := 0

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum -= node.Val
		if sum == 0 {
			cnt++
		}

		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}

	dfs(root, sum)

	return cnt +
		pathSum(root.Left, sum) +
		pathSum(root.Right, sum)
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
    def pathSum(self, root, sum):
        """
        :type root: TreeNode
        :type sum: int
        :rtype: int
        """
        def dfs(root, prevSum, sum):
            if not root:
                return 
            currSum = prevSum + root.val
            if currSum - sum in rec:
                self.count += rec[currSum - sum]
            if currSum in rec:
                rec[currSum] += 1
            else:
                rec[currSum] = 1
            dfs(root.left, currSum, sum)
            dfs(root.right, currSum, sum)
            rec[currSum] -= 1
            
        self.count = 0
        rec = {0:1}
        dfs(root, 0, sum)
        return self.count
```