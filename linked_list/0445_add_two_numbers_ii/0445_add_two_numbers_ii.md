445. Add Two Numbers II


Medium


You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Follow up:

What if you cannot modify the input lists? In other words, reversing the lists is not allowed.

Example:

```
Input: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 8 -> 0 -> 7
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // 把 l1 中的值，放入 stack s1
	s1 := make([]int, 0, 128)
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	// 把 l2 中的值，放入 stack s2
	s2 := make([]int, 0, 128)
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	sum := 0
	head := &ListNode{Val: 0}

	for len(s1) > 0 || len(s2) > 0 {
		// s1.pop()
		if len(s1) > 0 {
			sum += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}

		// s2.pop()
		if len(s2) > 0 {
			sum += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		head.Val = sum % 10
		ln := &ListNode{Val: sum / 10}
		ln.Next = head
		head = ln

		sum /= 10
	}

	if head.Val == 0 {
		return head.Next
	}
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
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        s1, s2 = [], []
        while l1:
            s1.append(l1.val)
            l1 = l1.next

        while l2:
            s2.append(l2.val)
            l2 = l2.next

        # tail = ListNode(-1)
        tail = ListNode(-1)
        # s
        s = 0
        while s1 or s2:
            a, b = s1.pop() if s1 else 0, s2.pop() if s2 else 0
            # +=
            s += a + b
            tail.val = s % 10
            s /= 10
            head = ListNode(s)
            head.next = tail
            tail = head

        # tail.val, or s
        return tail if tail.val else tail.next
```