328. Odd Even Linked List


Medium


Given a singly linked list, group all odd nodes together followed by the even nodes. Please note here we are talking about the node number and not the value in the nodes.

You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.

Example 1:

```
Input: 1->2->3->4->5->NULL
Output: 1->3->5->2->4->NULL
```

Example 2:

```
Input: 2->1->3->5->6->4->7->NULL
Output: 2->3->6->7->1->5->4->NULL
```

Constraints:

The relative order inside both the even and odd groups should remain as it was in the input.  
The first node is considered odd, the second node even and so on ...  
The length of the linked list is between [0, 10^4].  


## 方法

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	head2 := head.Next
	p1 := head
	p2 := head2
	for p1.Next != nil && p2.Next != nil {
		p1.Next = p2.Next
		p1 = p1.Next
		p2.Next = p1.Next
		p2 = p2.Next
	}
	p1.Next = head2
	return head
}
```


```python
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def oddEvenList(self, head: ListNode) -> ListNode:
        if not head:return head
        odd = head
        even_head = even = head.next
        while odd.next and even.next:
            odd.next = odd.next.next
            even.next = even.next.next
            odd,even = odd.next,even.next
        odd.next = even_head
        return head
```


