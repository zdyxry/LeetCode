class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def reorderList(self, head):
        """
        :type head: ListNode
        :rtype: None Do not return anything, modify head in-place instead.
        """
        if head == None or head.next == None:
            return
        
        mid = self.findMiddle(head)
        tail = self.reverse(mid.next)
        mid.next = None
        
        self.merge(head, tail)
        
    def reverse(self, head):
        newHead = None
        while head != None:
            temp = head.next
            head.next = newHead
            newHead = head
            head = temp
        return newHead
        
    def merge(self, head1, head2):
        index = 0
        dummy = ListNode(0)
        while head1 != None and head2 != None:
            if index % 2 == 0:
                dummy.next = head1
                head1 = head1.next
            else:
                dummy.next = head2
                head2 = head2.next
            index += 1
            dummy = dummy.next
        if head1:
            dummy.next = head1
        if head2:
            dummy.next = head2
    
    def findMiddle(self, head):
        slow, fast = head, head.next
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next
        return slow


head = ListNode(1)
head.next = ListNode(2)
head.next.next = ListNode(3)

Solution().reorderList(head)
print(head.next.val)