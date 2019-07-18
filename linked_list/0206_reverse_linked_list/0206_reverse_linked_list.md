206. Reverse Linked List


Easy


Reverse a singly linked list.

Example:
```
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
```

Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?


## 方法

1->2->3->4->5->NULL

head = 1
prev = NULL
temp = head.next                    temp = 2
head.next = prev                    1.next = NULL
prev = head                         prev = 1 -> NULL
head = temp                         head = 2

head = 2
prev = 1 -> NULL
temp = head.next                    temp = 3
head.next = prev                    2.next = 1 -> NULL
prev = head                         prev = 2 -> 1 -> NULL
head = temp                         head = 3

...
...
...


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode

	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}

	return prev
}
```


```python
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def reverseList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        prev = None
        while head:
            curr = head
            head = head.next
            curr.next = prev
            prev = curr
        return prev
```