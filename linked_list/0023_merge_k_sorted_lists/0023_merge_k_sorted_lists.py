
class ListNode:
    def __init__(self, val, next=None):
        self.val = val
        self.next = next


class Solution:
    def mergeKLists(self, lists: List[ListNode]) -> ListNode:
        if not lists or len(lists) == 0:
            return None
        all_vals = []
        for l in lists:
            while l:
                all_vals.append(l.val)
                l = l.next

        all_vals.sort()
        dummy = ListNode(None)
        cur = dummy
        for i in all_vals:
            temp_node = ListNode(i)
            cur.next = temp_node
            cur = temp_node
        return dummy.next