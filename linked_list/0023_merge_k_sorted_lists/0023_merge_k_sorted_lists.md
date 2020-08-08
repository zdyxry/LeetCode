23. Merge k Sorted Lists


Hard


Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

```
Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
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

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	guard := &ListNode{}
	for i := 0; i < len(lists); i++ {
		guard.Next = MergeTwoLists(guard.Next, lists[i])
	}

	return guard.Next
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	guard := &ListNode{}
	head := guard
	for {
		if l1 == nil {
			head.Next = l2
			break
		}
		if l2 == nil {
			head.Next = l1
			break
		}

		if l1.Val < l2.Val {
			v := l1.Val
			for l1 != nil && v == l1.Val {
				node := l1
				l1 = l1.Next
				head.Next = node
				head = head.Next
			}
		} else {
			v := l2.Val
			for l2 != nil && v == l2.Val {
				node := l2
				l2 = l2.Next
				head.Next = node
				head = head.Next
			}
		}
	}

	return guard.Next
}


```



```python
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
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
```

