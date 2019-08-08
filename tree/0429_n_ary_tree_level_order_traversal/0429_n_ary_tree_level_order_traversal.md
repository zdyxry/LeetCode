429. N-ary Tree Level Order Traversal
Easy

291

33

Favorite

Share
Given an n-ary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example, given a 3-ary tree:

![example](./narytreeexample.png)



We should return its level order traversal:
```
[
     [1],
     [3,2,4],
     [5,6]
]
```

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
    def levelOrder(self, root):
        """
        :type root: Node
        :rtype: List[List[int]]
        """
        if not root:return []
        res = []
        stack = [root]
        while stack:
            temp = []
            next_stack = []
            for node in stack:
                temp.append(node.val)
                for child in node.children:
                    next_stack.append(child)
            stack = next_stack
            res.append(temp)
        return res
```