class Node(object):
    def __init__(self, val=0, left=None, right=None, next=None):
        self.val = val
        self.left = left
        self.right = right
        self.next = next


class Solution(object):
    def connect1(self, root):
        if root is None:
            return
        
        the_root = root

        while root.left:
            next_layer = root.left
            while root.next:
                root.left.next = root.right
                root.right.next = root.next.left
                root = root.next
            root.left.next = root.right
            root = next_layer
        return the_root

    def connect2(self, root):
        from collections import deque
        q = deque()
        q.append(root)
        q.append(None)
        while root and q:
            node = q.popleft()
            if node:
                node.next = q[0]
                if node.left:
                    q.append(node.left)
                if node.right:
                    q.append(node.right)
            else:
                if len(q) > 0:
                    q.append(None)

        return root


root = Node(1)
root.left = Node(2)
root.left.left = Node(4)
root.left.right = Node(5)
root.right = Node(3)
root.right.left = Node(6)
root.right.right = Node(7)

res = Solution().connect1(root)
print(res.next)
