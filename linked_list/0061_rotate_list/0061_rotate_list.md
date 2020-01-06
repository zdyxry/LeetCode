61. Rotate List


Medium

Given a linked list, rotate the list to the right by k places, where k is non-negative.

Example 1:

```
Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL
```

Example 2:

```
Input: 0->1->2->NULL, k = 4
Output: 2->0->1->NULL
Explanation:
rotate 1 steps to the right: 2->0->1->NULL
rotate 2 steps to the right: 1->2->0->NULL
rotate 3 steps to the right: 0->1->2->NULL
rotate 4 steps to the right: 2->0->1->NULL
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
func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil {
		return head
	}

	tail := head
	for i := 0; i < k; i++ {
		if tail.Next == nil {
			// 处理 k 大于 list 的长度的情况
			// i+1 就是 list 的长度
			return rotateRight(head, k%(i+1))
		}
		tail = tail.Next
	}

	newTail := head
	for tail.Next != nil {
		newTail, tail = newTail.Next, tail.Next
	}

	newHead := newTail.Next
	newTail.Next, tail.Next = nil, head

	return newHead
}
```




```python
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def rotateRight(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        if not head or not head.next:
            return head

        n, cur = 1, head
        while cur.next:
            cur = cur.next
            n += 1
        cur.next = head

        cur, tail = head, cur
        for _ in xrange(n - k % n):
            tail = cur
            cur = cur.next
        tail.next = None

        return cur
```