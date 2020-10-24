1145. Binary Tree Coloring Game


Medium


Two players play a turn based game on a binary tree.  We are given the root of this binary tree, and the number of nodes n in the tree.  n is odd, and each node has a distinct value from 1 to n.

Initially, the first player names a value x with 1 <= x <= n, and the second player names a value y with 1 <= y <= n and y != x.  The first player colors the node with value x red, and the second player colors the node with value y blue.

Then, the players take turns starting with the first player.  In each turn, that player chooses a node of their color (red if player 1, blue if player 2) and colors an uncolored neighbor of the chosen node (either the left child, right child, or parent of the chosen node.)

If (and only if) a player cannot choose such a node in this way, they must pass their turn.  If both players pass their turn, the game ends, and the winner is the player that colored more nodes.

You are the second player.  If it is possible to choose such a y to ensure you win the game, return true.  If it is not possible, return false.

 

Example 1:


```
Input: root = [1,2,3,4,5,6,7,8,9,10,11], n = 11, x = 3
Output: true
Explanation: The second player can choose the node with value 2.
```

Constraints:

root is the root of a binary tree with n nodes and distinct node values from 1 to n.   
n is odd.   
1 <= x <= n <= 100


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

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {

    var (
        target *TreeNode
        count func(root *TreeNode) int
        findTarget func(root *TreeNode) 
    )

    count = func(root *TreeNode) int {
        if root == nil || root == target {
            return 0
        }
        return 1 + count(root.Left) + count(root.Right)
    }

    findTarget = func(root *TreeNode) {
        if root != nil {
           if root.Val == x {
                target = root
                return
            }
            findTarget(root.Left)
            findTarget(root.Right)
        }
    }

    findTarget(root)

    return count(root) > n/2 || count(target.Left) > n/2 || count(target.Right) > n/2
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
    def btreeGameWinningMove(self, root: TreeNode, n: int, x: int) -> bool:
        self.red_left, self.red_right = 0, 0
        
        def dfs(root):
            if not root:
                return 0
            left = dfs(root.left)
            right = dfs(root.right)
            if root.val == x:
                self.red_left = left
                self.red_right = right
            return left + right + 1
        
        dfs(root)
        parent = n - self.red_left - self.red_right - 1
        judge = [parent, self.red_left, self.red_right]
        return any([j > n // 2 for j in judge])

```