817. Linked List Components


Medium


We are given head, the head node of a linked list containing unique integer values.

We are also given the list G, a subset of the values in the linked list.

Return the number of connected components in G, where two values are connected if they appear consecutively in the linked list.

Example 1:

```
Input: 
head: 0->1->2->3
G = [0, 1, 3]
Output: 2
Explanation: 
0 and 1 are connected, so [0, 1] and [3] are the two connected components.
```

Example 2:

```
Input: 
head: 0->1->2->3->4
G = [0, 3, 1, 4]
Output: 2
Explanation: 
0 and 1 are connected, 3 and 4 are connected, so [0, 1] and [3, 4] are the two connected components.
```

Note:

If N is the length of the linked list given by head, 1 <= N <= 10000.  
The value of each node in the linked list will be in the range [0, N - 1].  
1 <= G.length <= 10000.  
G is a subset of all values in the linked list.


## 方法

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, G []int) int {
    m := make(map[int]bool)
	for i := range G {
		m[G[i]] = true
	}

	connected := m[head.Val]

	cnt := 0
	for head != nil {
		if connected != m[head.Val] {
			if connected == false {
				connected = true
			} else {
				connected = false
				cnt++
			}
		}
		head = head.Next
	}
	
	if connected{
		return cnt+1
	}
	return cnt
}
```



```python
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def numComponents(self, head: ListNode, G: List[int]) -> int:
        s=set(G)
        
        currentNode=head
        stack=[]
        count=0
        while(currentNode):
            v=currentNode.val
            if v not in s:
                if stack:
                    count+=1
                    stack=[]
            if v in s:
                stack.append(v)
            currentNode=currentNode.next
        if stack:
            count+=1
        return count
```