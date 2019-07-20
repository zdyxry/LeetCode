234. Palindrome Linked List


Easy


Given a singly linked list, determine if it is a palindrome.

Example 1:
```
Input: 1->2
Output: false
```

Example 2:
```
Input: 1->2->2->1
Output: true
```

Follow up:

Could you do it in O(n) time and O(1) space?


## 方法


```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
		return true
	}
	res := true
	// 寻找中间结点
	p1 := head
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
    fmt.Println(p1, p2)
    
	// 反转链表后半部分  1->2->3->4->5->6 to 1->2->3->6->5->4
	preMiddle := p1    //3
	preCurrent := p1.Next  // 4
    fmt.Println(preMiddle, preCurrent)
	for preCurrent.Next != nil {
		current := preCurrent.Next
        fmt.Println("current :", current)
		preCurrent.Next = current.Next
        fmt.Println("preCurrent.Next :", preCurrent.Next)
		current.Next = preMiddle.Next
        fmt.Println("current.Next :", current.Next)
		preMiddle.Next = current
        fmt.Println("preMiddle.Next :", preMiddle.Next)
	}

	// 扫描表，判断是否是回文
	p1 = head
	p2 = preMiddle.Next
	for p1 != preMiddle {
		if p1.Val == p2.Val {
			p1 = p1.Next
			p2 = p2.Next
		} else {
			res = false
			break
		}
	}
	if p1 == preMiddle {
		if p2 != nil && p1.Val != p2.Val {
			return false
		}
	}

	return res
}

```


```python
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def isPalindrome(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        fast = slow = head
        # find the mid node
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next
        # reverse the second half
        node = None
        while slow:
            nxt = slow.next
            slow.next = node
            node = slow
            slow = nxt
        # compare the first and second half nodes
        while node: # while node and head:
            if node.val != head.val:
                return False
            node = node.next
            head = head.next
        return True
```