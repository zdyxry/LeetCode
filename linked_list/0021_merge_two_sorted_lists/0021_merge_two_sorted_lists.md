21. Merge Two Sorted Lists


Easy

Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:
```
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4

```


## 方法
定义链表，遍历链表比较当前值，若当前值较小，则将其追加到结果中，并重置链表。


```go
type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }

    var head, node *ListNode
    if l1.Val < l2.Val {
        head = l1
        node = l1
        l1 = l1.Next
    } else {
        head = l2
        node = l2
        l2 = l2.Next
    }

    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            node.Next = l1
            l1 = l1.Next
        } else {
            node.Next = 12
            l2 = l2.Next
        }
        node = node.Next
    }

    if l1 != nil {
        node.Next = l1
    }

    if l2 != nil {
        node.Next = l2
    }

    return head 
}

```

```python
# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        curr = dummy = ListNode(0)
        while l1 and l2:
            if l1.val < l2.val:
                curr.next = l1
                l1 = l1.next
            else:
                curr.next = l2
                l2 = l2.next
            curr = curr.next
        curr.next = l1 or l2
        return dummy.next
```