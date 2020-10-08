173. Binary Search Tree Iterator


Medium


Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.

Calling next() will return the next smallest number in the BST.

 

Example:



```
BSTIterator iterator = new BSTIterator(root);
iterator.next();    // return 3
iterator.next();    // return 7
iterator.hasNext(); // return true
iterator.next();    // return 9
iterator.hasNext(); // return true
iterator.next();    // return 15
iterator.hasNext(); // return true
iterator.next();    // return 20
iterator.hasNext(); // return false
```

Note:

next() and hasNext() should run in average O(1) time and uses O(h) memory, where h is the height of the tree.  
You may assume that next() call will always be valid, that is, there will be at least a next smallest number in the BST when next() is called.   


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
type BSTIterator struct {
    nums []int
    root *TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    nums := make([]int,0)
    inorder(root,&nums)
    return BSTIterator{root:root,nums:nums}
}

func inorder(root *TreeNode,nums *[]int){
    if root == nil{
        return
    }
    inorder(root.Left,nums)
    *nums = append(*nums,root.Val)
    inorder(root.Right,nums)
}
/** @return the next smallest number */
func (this *BSTIterator) Next() int {
    res:=this.nums[0]
    this.nums = this.nums[1:]
    return res
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
    if len(this.nums) > 0{
        return true
    }
    return false
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

```


```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class BSTIterator:

    def __init__(self, root: TreeNode):
        self.index = 0
        self.res = []
        self.dfs(root)

    def next(self) -> int:
        """
        @return the next smallest number
        """
        tmp = self.index
        self.index += 1
        return self.res[tmp]

    def hasNext(self) -> bool:
        """
        @return whether we have a next smallest number
        """
        if self.index == len(self.res):
            return False
        return True

    def dfs(self, root):
        if not root:
            return
        self.dfs(root.left)
        self.res.append(root.val)
        self.dfs(root.right)
```