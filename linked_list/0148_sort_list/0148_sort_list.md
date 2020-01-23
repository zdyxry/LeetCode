148. Sort List

Medium


Sort a linked list in O(n log n) time using constant space complexity.

Example 1:

```
Input: 4->2->1->3
Output: 1->2->3->4
```

Example 2:

```
Input: -1->5->3->4->0
Output: -1->0->3->4->5
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
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	left, right := split(head)
	return merge(sortList(left), sortList(right))
}

// 从中间位置，切分 list
func split(head *ListNode) (left, right *ListNode) {
	// head.Next != nil
	// 因为， sortList 已经帮忙检查过了

	// fast 的变化速度是 slow 的两倍
	// 当 fast 到达末尾的时候，slow 正好在 list 的中间
	slow, fast := head, head
	var slowPre *ListNode

	for fast != nil && fast.Next != nil {
		slowPre, slow = slow, slow.Next
		fast = fast.Next.Next
	}

	// 斩断 list
	slowPre.Next = nil
	// 使用 slowPre 是为了确保当 list 的长度为 2 时，会均分为两个长度为 1 的子 list

	left, right = head, slow
	return
}

// 把已经排序好了的两个 list left 和 right
// 进行合并
func merge(left, right *ListNode) *ListNode {
	// left != nil , right != nil
	// 因为， sortList 已经帮忙检查过了

	cur := &ListNode{}
	headPre := cur
	for left != nil && right != nil {
		// cur 总是接上较小的 node
		if left.Val < right.Val {
			cur.Next, left = left, left.Next
		} else {
			cur.Next, right = right, right.Next
		}
		cur = cur.Next
	}

	// 处理 left 或 right 中，剩下的节点
	if left == nil {
		cur.Next = right
	} else {
		cur.Next = left
	}

	return headPre.Next
}
```



```python
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def sortList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        myhead = head
        mylist = []
        while head:
            mylist.append(head.val)
            head = head.next
        mylist.sort()
        head = myhead
        for i in range(len(mylist)):
            head.val = mylist[i]
            head = head.next
        return myhead
```


```python
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def sortList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if head == None or head.next == None:
            return head

        fast, slow, prev = head, head, None
        while fast != None and fast.next != None:
            prev, fast, slow = slow, fast.next.next, slow.next
        :q
        .next = None

        sorted_l1 = self.sortList(head)
        sorted_l2 = self.sortList(slow)

        return self.mergeTwoLists(sorted_l1, sorted_l2)

    def mergeTwoLists(self, l1, l2):
        dummy = ListNode(0)

        cur = dummy
        while l1 != None and l2 != None:
            if l1.val <= l2.val:
                cur.next, cur, l1 = l1, l1, l1.next
            else:
                cur.next, cur, l2 = l2, l2, l2.next

        if l1 != None:
            cur.next = l1
        if l2 != None:
            cur.next = l2

        return dummy.next
```
