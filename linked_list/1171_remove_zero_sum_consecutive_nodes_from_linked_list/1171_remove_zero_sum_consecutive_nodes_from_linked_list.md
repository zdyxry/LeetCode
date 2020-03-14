1171. Remove Zero Sum Consecutive Nodes from Linked List


Medium


Given the head of a linked list, we repeatedly delete consecutive sequences of nodes that sum to 0 until there are no such sequences.

After doing so, return the head of the final linked list.  You may return any such answer.

 

(Note that in the examples below, all sequences are serializations of ListNode objects.)

Example 1:

```
Input: head = [1,2,-3,3,1]
Output: [3,1]
Note: The answer [1,2,1] would also be accepted.
```

Example 2:

```
Input: head = [1,2,3,-3,4]
Output: [1,2,4]
```

Example 3:

```
Input: head = [1,2,3,-3,-2]
Output: [1]
```

 

Constraints:

The given linked list will contain between 1 and 1000 nodes.  
Each node in the linked list has -1000 <= node.val <= 1000.

## 方法

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
    // preSum Map存放最新sum的结点(*******sum相同存放后一个结点，利用这个特效，如果当前的sum有对应的值的话，直接让当前node.next = map[sum].next, 如果全部分sum = 后部分sum, 说明中间区间sum=0, 那么直接让之前结点.next = 后部分结点开始.next；如果只有这一个sum,node.next = map[sum].next 还是原结点.next=原结点.next
    preSumMap := make(map[int]*ListNode)
    
    sum := 0
    dummy := &ListNode{0, nil}
    dummy.Next = head
    cur := dummy
    
    for ; cur != nil; cur = cur.Next {
        sum += cur.Val
        preSumMap[sum] = cur
    }
    
    sum = 0
    cur = dummy
    for ; cur != nil; cur = cur.Next {
        sum += cur.Val
        cur.Next = preSumMap[sum].Next
    }
    
    return dummy.Next
}
```


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
    type NodeSum struct {
        Node *ListNode
        Sum int
    }
    acc, sum, dummy := 0, make(map[int]NodeSum), &ListNode{Next: head}
    sum[0] = NodeSum{dummy, 0}
    for curr := head; curr != nil; curr = curr.Next {
        acc += curr.Val
        if p, ok := sum[acc]; ok {
            for node, subSum := p.Node.Next, p.Sum; node != curr; node = node.Next {
                subSum += node.Val
                delete(sum, subSum)
            }
            p.Node.Next = curr.Next
        } else {
            sum[acc] = NodeSum{curr, acc}
        }
    }
    return dummy.Next
}

```



```python
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def removeZeroSumSublists(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        dummy = ListNode(0)
        dummy.next = head
        prefix = 0
        d = {0:dummy} # key is the prefix sum, value is the last node of getting this sum value, which is l5
        while head:
            prefix += head.val
            d[prefix] = head
            head = head.next
		# Go from the dummy node again to set the next node to be the last node for a prefix sum 
        head = dummy
        prefix = 0
        while head:
            prefix += head.val
            head.next = d[prefix].next
            head = head.next
        
        return dummy.next
```