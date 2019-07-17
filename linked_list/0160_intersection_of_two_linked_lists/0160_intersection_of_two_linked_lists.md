160. Intersection of Two Linked Lists


Easy


Write a program to find the node at which the intersection of two singly linked lists begins.

For example, the following two linked lists:


begin to intersect at node c1.

 

Example 1:

```
Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
Output: Reference of the node with value = 8
Input Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,0,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.
```

Example 2:

```
Input: intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
Output: Reference of the node with value = 2
Input Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [0,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.
```

Example 3:

```
Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
Output: null
Input Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
Explanation: The two lists do not intersect, so return null.
```

Notes:

If the two linked lists have no intersection at all, return null.

The linked lists must retain their original structure after the function returns.

You may assume there are no cycles anywhere in the entire linked structure.

Your code should preferably run in O(n) time and use only O(1) memory.

## 方法

当 A 链表循环完成后，循环 B 链表；
当 B 链表循环完成后，循环 A 链表。

不会陷入死循环的原因：当 A 与 B 两个链表最终走完的路程均为 A+B，最终值为 nil，会退出循环。

A: 1 -> 2 -> 3
B: 4 -> 5

比较顺序如下：

1 vs 4 
2 vs 5
3 vs nil
nil vs 1
4 vs 2
5 vs 3 
nil vs nil

A: 1 -> 2 -> 3
B: 2 -> 3

比较顺序如下：

1 vs 2 
2 vs 3 
3 vs nil
nil vs 1
2 vs 2 
3 vs 3 



```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    a, b := headA, headB
	hasLinkedToB, hasLinkedToA := false, false
	for a != nil && b != nil {
		if a == b {
			return b
		}
		a, b = a.Next, b.Next
		if a == nil && !hasLinkedToB {
			a = headB
			hasLinkedToB = true
		}
		if b == nil && !hasLinkedToA {
			b = headA
			hasLinkedToA = true
		}
	}
	return nil
}
```


```python
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def getIntersectionNode(self, headA, headB):
        """
        :type head1, head1: ListNode
        :rtype: ListNode
        """
        if not headA or not headB:
            return None

        p,q = headA , headB
        while p and q and p != q :
            p = p.next
            q = q.next
            if p == q :
                return q
            if not p :
                p = headB
            if not q :
                q = headA
        return p

```