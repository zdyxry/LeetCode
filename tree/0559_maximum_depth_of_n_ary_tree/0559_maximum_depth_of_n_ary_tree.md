559. Maximum Depth of N-ary Tree
Easy

413

22

Favorite

Share
Given a n-ary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

For example, given a 3-ary tree:

![example](./559.png)



We should return its max depth, which is 3.

 

Note:

The depth of the tree is at most 1000.

The total number of nodes is at most 5000.


## 方法



```python
"""
# Definition for a Node.
class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children
"""
class Solution(object):
    def maxDepth(self, root):
        """
        :type root: Node
        :rtype: int
        """
        if not root:
            return 0
        queue = []
        queue.append(root)
        height = 0
        while len(queue)>0:
            num_nodes = len(queue)
            for i in range(num_nodes):
                for child in queue[0].children:
                    queue.append(child)
                queue.remove(queue[0])
            height+=1
        return height
```