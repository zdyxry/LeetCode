1315. Sum of Nodes with Even-Valued Grandparent


Medium


Given a binary tree, return the sum of values of nodes with even-valued grandparent.  (A grandparent of a node is the parent of its parent, if it exists.)

If there are no nodes with an even-valued grandparent, return 0.

 

Example 1:



```
Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
Output: 18
Explanation: The red nodes are the nodes with even-value grandparent while the blue nodes are the even-value grandparents.
```


Constraints:

The number of nodes in the tree is between 1 and 10^4.

The value of nodes is between 1 and 100.


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
func sumEvenGrandparent(root *TreeNode) int {
    if root == nil {
		return 0
	}
	return sumEvenGrandparent2(root.Left, root.Val % 2 == 0, false) +
		sumEvenGrandparent2(root.Right, root.Val % 2 == 0, false)
}

func sumEvenGrandparent2(node *TreeNode, isParentEven bool, isGrandParentEven bool) int {
	sum := 0
	if node == nil {
		return sum
	}

	if isGrandParentEven {
		sum += node.Val
	}

	return sum +
		sumEvenGrandparent2(node.Left, node.Val % 2 == 0, isParentEven) +
		sumEvenGrandparent2(node.Right, node.Val % 2 == 0, isParentEven)
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
func sumEvenGrandparent(root *TreeNode) int {
    // Handle leaf nodes
    if root == nil{
        return 0
    }
    
    sum := 0
    // Condition for even-value root
    if root.Val % 2 == 0{
        // Check presence of grand-nodes
        if root.Right != nil{
            if root.Right.Left != nil{
                sum += root.Right.Left.Val
            }
            if root.Right.Right != nil{
                sum += root.Right.Right.Val
            }
        }
        if root.Left != nil{
            if root.Left.Left != nil{
                sum += root.Left.Left.Val
            }
            if root.Left.Right != nil{
                sum += root.Left.Right.Val
            }
        }
        }
    
    // Iteration logic
    sum += sumEvenGrandparent(root.Left)
    sum += sumEvenGrandparent(root.Right)
    
    return sum
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
    def sumEvenGrandparent(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        grandfa = father = None
        res = 0
        deque = collections.deque()
        deque.append((root, father, grandfa))
        while deque:
            node, father, grandfa = deque.popleft()
            if node:
                if grandfa and grandfa.val % 2 == 0:
                    res += node.val
                deque.append((node.left, node, father))
                deque.append((node.right, node, father))

        return res
```