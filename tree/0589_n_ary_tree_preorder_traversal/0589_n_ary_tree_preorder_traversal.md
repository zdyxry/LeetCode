589. N-ary Tree Preorder Traversal
Easy

278

38

Favorite

Share
Given an n-ary tree, return the preorder traversal of its nodes' values.

For example, given a 3-ary tree:

```
            1
          / | \
         3  2  4
        / \
       5   6
```



 

Return its preorder traversal as: [1,3,5,6,2,4].

 

Note:

Recursive solution is trivial, could you do it iteratively?


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
    def preorder(self, root):
        """
        :type root: Node
        :rtype: List[int]
        """
        result = []
        
        def helper(node):
            if node is None:
                return
            #Leaf node
            result.append(node.val)
            for child in node.children:
                helper(child)
        
        helper(root)
        return result
```



```python
"""
# Definition for a Node.
class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children
"""
class Solution(object):
    def preorder(self, root):
        """
        :type root: Node
        :rtype: List[int]
        """
        if not root:
            return []
        traversal = []
        stack = [root]
        while stack:
            cur = stack.pop()
            traversal.append(cur.val)
            stack.extend(reversed(cur.children))
        return traversal
```