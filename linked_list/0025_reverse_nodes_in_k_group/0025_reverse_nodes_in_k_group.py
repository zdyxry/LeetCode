
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        curr = head
        for _ in range(k):
            if not curr: return head
            curr = curr.next

        prev = None
        curr = head
        for _ in range(k):
            nxt = curr.next
            curr.next = prev
            prev = curr
            curr = nxt

        head.next = self.reverseKGroup(curr, k)
        return prev


head = ListNode(0)
cur = head
for i in range(1,5):
    head.next = ListNode(i)
    head = head.next

head = cur
while head:
    print(head.val)
    head = head.next
print("###################")
res = Solution().reverseKGroup(cur, 3)
head = res
while head:
    print(head.val)
    head = head.next