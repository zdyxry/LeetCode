1721. Swapping Nodes in a Linked List


Medium


You are given the head of a linked list, and an integer k.

Return the head of the linked list after swapping the values of the kth node from the beginning and the kth node from the end (the list is 1-indexed).

 

Example 1:

```
Input: head = [1,2,3,4,5], k = 2
Output: [1,4,3,2,5]
```

Example 2:

```
Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
Output: [7,9,6,6,8,7,3,0,9,5]
```

Example 3:

```
Input: head = [1], k = 1
Output: [1]
```

Example 4:

```
Input: head = [1,2], k = 1
Output: [2,1]
```

Example 5:

```
Input: head = [1,2,3], k = 2
Output: [1,2,3]
```
 

Constraints:

The number of nodes in the list is n.   
1 <= k <= n <= 105   
0 <= Node.val <= 100


## 方法

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
    slow, fast := head, head
    for i:=0;i<k-1;i++ {
        fast = fast.Next
    }
    left := fast
    for fast.Next != nil {
        slow, fast = slow.Next, fast.Next
    }
    slow.Val, left.Val = left.Val, slow.Val
    return head
}
```


```python
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def swapNodes(self, head: ListNode, k: int) -> ListNode:
        slow = fast = tmp = head
        for _ in range(k-1):
            fast, tmp = fast.next, tmp.next
        while fast.next:
            fast = fast.next
            slow = slow.next
        tmp.val, slow.val = slow.val, tmp.val
        return head
```