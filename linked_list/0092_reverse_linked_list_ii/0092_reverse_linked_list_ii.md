92. Reverse Linked List II


Medium


Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

```
Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
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
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if head == nil || m >= n {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead
	for count := 0; pre.Next != nil && count < m-1; count++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next
	for i := 0; i < n-m; i++ {
		tmp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = tmp
	}
	return newHead.Next
}
```



```python
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def reverseBetween(self, head, m, n):
        """
        :type head: ListNode
        :type m: int
        :type n: int
        :rtype: ListNode
        """
        diff, dummy, cur = n - m + 1, ListNode(-1), head
        dummy.next = head

        last_unswapped = dummy
        while cur and m > 1:
            cur, last_unswapped, m = cur.next, cur, m - 1

        prev, first_swapped = last_unswapped,  cur
        while cur and diff > 0:
            cur.next, prev, cur, diff = prev, cur, cur.next, diff - 1

        last_unswapped.next, first_swapped.next = prev, cur

        return dummy.next
```