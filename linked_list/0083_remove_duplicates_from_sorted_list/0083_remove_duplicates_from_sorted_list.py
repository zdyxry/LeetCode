class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def deleteDuplicates(self, head):
        if not head or not head.next:
            return head

        prev = head
        curr = head.next

        while curr:
            if prev.val == curr.val:
                prev.next = prev.next.next
            else:
                prev = curr
            curr = curr.next
        return head


head = ListNode(1)
head.next = ListNode(1)
head.next.next = ListNode(2)

result = Solution().deleteDuplicates(head)
print(result.next.val)