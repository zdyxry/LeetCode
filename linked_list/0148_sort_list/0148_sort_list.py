
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def sortList(self, head):
        myhead = head
        mylist = []
        while head:
            mylist.append(head.val)
            head = head.next
        mylist.sort()
        head = myhead
        for i in range(len(mylist)):
            head.val = mylist[i]
            head = head.next
        return myhead

    def sortList2(self, head):
        if head == None or head.next == None:
            return head
        
        fast, slow, prev = head, head, None
        while fast != None and fast.next != None:
            prev, fast, slow = slow, fast.next.next, slow.next
        prev.next = None

        sorted_l1 = self.sortList2(head)
        sorted_l2 = self.sortList2(slow)

        return self.mergeTwoList(sorted_l1, sorted_l2)

    def mergeTwoList(self, l1, l2):
        dummy = ListNode(0)

        cur = dummy
        while l1 != None and l2 != None:
            if l1.val <= l2.val:
                cur.next, cur, l1 = l1, l1, l1.next
            else:
                cur.next, cur, l2 = l2, l2, l2.next
        
        if l1 != None:
            cur.next = l1
        if l2 != None:
            cur.next = l2

        return dummy.next

head = ListNode(3)
head.next = ListNode(2)
head.next.next = ListNode(1)
res = Solution().sortList2(head)
print(res.val)