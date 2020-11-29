1669. Merge In Between Linked Lists


Medium

You are given two linked lists: list1 and list2 of sizes n and m respectively.

Remove list1's nodes from the ath node to the bth node, and put list2 in their place.

The blue edges and nodes in the following figure incidate the result:


Build the result list and return its head.

 

Example 1:


```
Input: list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
Output: [0,1,2,1000000,1000001,1000002,5]
Explanation: We remove the nodes 3 and 4 and put the entire list2 in their place. The blue edges and nodes in the above figure indicate the result.
```

Example 2:


```
Input: list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1000003,1000004]
Output: [0,1,1000000,1000001,1000002,1000003,1000004,6]
Explanation: The blue edges and nodes in the above figure indicate the result.
```

Constraints:

```
3 <= list1.length <= 104
1 <= a <= b < list1.length - 1
1 <= list2.length <= 104
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
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    n := list1
    var startRef, endRef *ListNode
    for i := 0; i <= b; i++ {
        if i == a-1 {
            startRef = n
        }
        
        if i == b {
            endRef = n
        }
        
        n = n.Next
    }
    
    startRef.Next = list2
    n = list2
    for n.Next != nil {
        n = n.Next
    }
    
    n.Next = endRef.Next
    return list1
}
```


```python
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeInBetween(self, list1: ListNode, a: int, b: int, list2: ListNode) -> ListNode:
        dummy = ListNode(0)
        dummy.next = list1
        l = dummy
        c = -1
        d1,d2 = None,None
        while l:
            if c == a - 1:
                d1 = l
            if c == b + 1:
                d2 = l
            l = l.next
            c += 1
        if d2 == None:
            d1.next = list2
        else:
            d1.next = list2
            q = list2
            while q.next:
                q = q.next
            q.next = d2
        return dummy.next
```