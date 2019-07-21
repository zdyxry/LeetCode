class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def deleteNode(self, Node):
        node.val = node.next.val
        node.next = node.next.next


node = ListNode(1)
node.next = ListNode(2)
node.next.next = ListNode(3)
Solution().deleteNode(node.next)

print(node.next.val)