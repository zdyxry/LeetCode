1305. All Elements in Two Binary Search Trees


Medium


Given two binary search trees root1 and root2.

Return a list containing all the integers from both trees sorted in ascending order.

 

Example 1:


```
Input: root1 = [2,1,4], root2 = [1,0,3]
Output: [0,1,1,2,3,4]
```

Example 2:

```
Input: root1 = [0,-10,10], root2 = [5,1,7,0,2]
Output: [-10,0,0,1,2,5,7,10]
```

Example 3:

```
Input: root1 = [], root2 = [5,1,7,0,2]
Output: [0,1,2,5,7]
```

Example 4:

```
Input: root1 = [0,-10,10], root2 = []
Output: [-10,0,10]
```

Example 5:


```
Input: root1 = [1,null,8], root2 = [8,1]
Output: [1,1,8,8]
```
 

Constraints:

Each tree has at most 5000 nodes.   
Each node's value is between [-10^5, 10^5].


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
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    res := make([]int,0,10)
    getList(root1,&res)
    getList(root2,&res)
    sort.Ints(res)
    return res
}
func getList(root *TreeNode,res *[]int){
    if root == nil{
        return
    }
    *res = append(*res,root.Val)
    getList(root.Left,res)
    getList(root.Right,res)
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
    def getAllElements(self, root1: TreeNode, root2: TreeNode) -> List[int]:
        ans = list()

        def dfs(node):
            if not node:
                return
            ans.append(node.val)
            dfs(node.left)
            dfs(node.right)
        
        dfs(root1)
        dfs(root2)
        ans.sort()
        return ans
```