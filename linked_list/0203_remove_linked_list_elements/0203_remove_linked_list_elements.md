203. Remove Linked List Elements


Easy


Remove all elements from a linked list of integers that have value val.

Example:
```
Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
```


## 方法




```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    dummy := new(ListNode)
    dummy.Next = head    
    p := dummy
    
    for p.Next != nil {
        if (p.Next.Val == val) {
            p.Next = p.Next.Next
        } else {
            p = p.Next
        }
    }
    
    return dummy.Next
}
```



```python
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def removeElements(self, head, val):
        """
        :type head: ListNode
        :type val: int
        :rtype: ListNode
        """
        dummy = ListNode(0)
        dummy.next = head
        pre = dummy
        while head:
            if head.val == val:
                pre.next, head = head.next, head.next
            else:
                pre, head = head, head.next
        return dummy.next
```