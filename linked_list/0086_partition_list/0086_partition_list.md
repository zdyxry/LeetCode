86. Partition List


Medium


Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Example:

```
Input: head = 1->4->3->2->5->2, x = 3
Output: 1->2->2->4->3->5
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
func partition(head *ListNode, x int) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

	// 存放 <x 的节点的链
	lessHead := &ListNode{}
	// 存放 >=x 的节点的链
	noLessHead := &ListNode{}
	// Head.Next 才是真正的 head
	// 这样处理是为了 for 循环中的逻辑简单

	lessEnd := lessHead
	noLessEnd := noLessHead

	for head != nil {
		if head.Val < x {
			lessEnd.Next = head
			lessEnd = lessEnd.Next
		} else {
			noLessEnd.Next = head
			noLessEnd = noLessEnd.Next
		}
		head = head.Next
	}

	// 把两部分首尾相连
	lessEnd.Next = noLessHead.Next
	// 注意封闭 noLessEnd
	noLessEnd.Next = nil

	head = lessHead.Next
	return head
}
```



```python
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def partition(self, head, x):
        """
        :type head: ListNode
        :type x: int
        :rtype: ListNode
        """
        dummySmaller, dummyGreater = ListNode(-1), ListNode(-1)
        smaller, greater = dummySmaller, dummyGreater

        while head:
            if head.val < x:
                smaller.next = head
                smaller = smaller.next
            else:
                greater.next = head
                greater = greater.next
            head = head.next

        smaller.next = dummyGreater.next
        greater.next = None

        return dummySmaller.next
```