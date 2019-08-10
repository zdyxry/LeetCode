501. Find Mode in Binary Search Tree


Easy


Given a binary search tree (BST) with duplicates, find all the mode(s) (the most frequently occurred element) in the given BST.

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than or equal to the node's key.
The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
Both the left and right subtrees must also be binary search trees.
 

For example:
```
Given BST [1,null,2,2],

   1
    \
     2
    /
   2
 

return [2].
```

Note: If a tree has more than one mode, you can return them in any order.


Follow up: Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).


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
func findMode(root *TreeNode) []int {
    if root==nil {return nil}
    
    prev := -1<<31
    count := 0
    max := 1
    result := make([]int,0)
    findBSTMode(root, &prev,&count,&max, &result)
    
    return result
}

func findBSTMode(root *TreeNode, prev,count,max *int, result *[]int) {
    if root==nil {return}
    
    // Go Left
    findBSTMode(root.Left, prev, count, max, result)
    
    // Set count,max and previous:
    if root.Val==*prev {*count++} else {*count=1}
    if *max < *count { // Found a better max
        *result = []int{root.Val} // Reset result array
        *max = *count // Set max
    } else if *max == *count {
        *result = append(*result, root.Val) // Append to result
    }
    *prev = root.Val
    
    // Go Right
    findBSTMode(root.Right, prev, count, max, result)
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
func findMode(root *TreeNode) []int {
    r := map[int]int{}
	search(root, r)

	max := -1
	res := []int{}
	for n, v := range r {
		if max <= v {
			if max < v {
				max = v
				res = res[0:0]
			}
			res = append(res, n)
		}
	}

	return res
}

func search(root *TreeNode, rec map[int]int) {
	if root == nil {
		return
	}

	rec[root.Val]++

	search(root.Left, rec)
	search(root.Right, rec)
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
    def findMode(self, root):
        """
        :type root: TreeNode
        :rtype: List[int]
        """
        if not root:
            return []
        
        self.count = 0
        self.mode_count = 0
        self.modes = []
        self.last_seen = None
        
        def helper(node):
            
            if node.left:
                helper(node.left)
                
            if node.val != self.last_seen:
                if self.count > self.mode_count:
                    self.mode_count = self.count
                    self.modes = [self.last_seen]
                elif self.count == self.mode_count:
                    self.modes.append(self.last_seen)
                self.last_seen = node.val
                self.count = 1
            else:
                self.count += 1
                
            if node.right:
                helper(node.right)
        
        helper(root)
        if self.count > self.mode_count:
            return [self.last_seen]
        elif self.count == self.mode_count:
            self.modes.append(self.last_seen)
            return self.modes
        else:
            return self.modes
```