import collections

class TreeNode(object):
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution(object):
    def sumEvenGrandparent(self, root):
        grandfa = father = None
        res = 0
        deque = collections.deque()
        deque.append((root, father, grandfa))
        while deque:
            node, father, grandfa = deque.popleft()
            if node:
                if grandfa and grandfa % 2 == 0:
                    res += node.val
                deque.append((node.left, node, father))
                deque.append((node.right, node, father))
