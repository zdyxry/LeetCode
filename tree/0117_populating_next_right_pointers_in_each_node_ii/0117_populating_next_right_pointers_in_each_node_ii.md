117. Populating Next Right Pointers in Each Node II


Medium


Given a binary tree

```
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
```

Initially, all next pointers are set to NULL.

 

Follow up:

You may only use constant extra space.

Recursive approach is fine, you may assume implicit stack space does not count as extra space for this problem.
 

Example 1:

![0117](0117.png)

```
Input: root = [1,2,3,4,5,null,7]
Output: [1,#,2,3,#,4,5,7,#]
Explanation: Given the above binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.
```

Constraints:

The number of nodes in the given tree is less than 6000.

-100 <= node.val <= 100

## 方法

```python
"""
# Definition for a Node.
class Node(object):
    def __init__(self, val=0, left=None, right=None, next=None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next
"""
class Solution(object):
    def connect(self, root):
        """
        :type root: Node
        :rtype: Node
        """
        old_root = root
        prekid = Node(0)
        kid = prekid   # Let kid point to prekid 
        while root:
            while root:
                if root.left:
                    kid.next = root.left
                    kid = kid.next
                if root.right:
                    kid.next = root.right
                    kid = kid.next
                root = root.next
            root, kid = prekid.next, prekid
            kid.next = None  # Reset the chain for prekid
        return old_root
```