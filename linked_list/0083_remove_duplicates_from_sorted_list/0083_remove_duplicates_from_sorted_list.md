83. Remove Duplicates from Sorted List


Easy


Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:
```
Input: 1->1->2
Output: 1->2
```

Example 2:
```
Input: 1->1->2->3->3
Output: 1->2->3
```

## 方法

考虑边界情况，当 head 为空或 head.next 为空时，直接返回。

如果当前值与下一个节点的数值相同，则跳过下一个节点。



```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
		return nil
	}

	temp := head

	for temp.Next != nil {
		if temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}
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
    def deleteDuplicates(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if not head or not head.next:
            return head
        
        prev = head
        curr = head.next
        
        while curr:
            if prev.val == curr.val:
                prev.next = prev.next.next
            
            else:
                prev = curr
            
            curr = curr.next
        
        return head

```