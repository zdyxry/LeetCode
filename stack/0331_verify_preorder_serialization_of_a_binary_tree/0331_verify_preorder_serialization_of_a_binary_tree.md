331. Verify Preorder Serialization of a Binary Tree


Medium


One way to serialize a binary tree is to use pre-order traversal. When we encounter a non-null node, we record the node's value. If it is a null node, we record using a sentinel value such as #.

```
     _9_
    /   \
   3     2
  / \   / \
 4   1  #  6
/ \ / \   / \
# # # #   # #
```

For example, the above binary tree can be serialized to the string "9,3,4,#,#,1,#,#,2,#,6,#,#", where # represents a null node.

Given a string of comma separated values, verify whether it is a correct preorder traversal serialization of a binary tree. Find an algorithm without reconstructing the tree.

Each comma separated value in the string must be either an integer or a character '#' representing null pointer.

You may assume that the input format is always valid, for example it could never contain two consecutive commas such as "1,,3".

Example 1:

```
Input: "9,3,4,#,#,1,#,#,2,#,6,#,#"
Output: true
```

Example 2:

```
Input: "1,#"
Output: false
```

Example 3:

```
Input: "9,#,#,1"
Output: false
```

## 方法


```go
// 用一个栈，从字符串的左侧向右依次进栈，如果满足栈的后三位是数字，#，#的模式时，
// 说明可以构成合法的叶子节点，把这个叶子节点换成#号，代表空节点，
// 然后继续遍历。最后应该只剩下一个#，那么就是一个合法的二叉树。
func isValidSerialization(preorder string) bool {
    parts := strings.Split(preorder, ",")
	s := make([]string, 0, len(parts))
	for _, p := range parts {
		for p == "#" && len(s) > 0 && s[len(s)-1] == "#" {
			s = s[:len(s)-1]
			if len(s) == 0 {
				return false
			}
			s = s[:len(s)-1]
		}
		s = append(s, p)
	}
	return len(s) == 1 && s[0] == "#"
}

```


```go
func isValidSerialization(preorder string) bool {
    nodes, diff := strings.Split(preorder, ","), 1
	for _, node := range nodes {
		diff--
		if diff < 0 {
			return false
		}
		if node != "#" {
			diff += 2
		}
	}
	return diff == 0
}
```



```python
class Solution(object):
    def isValidSerialization(self, preorder):
        """
        :type preorder: str
        :rtype: bool
        """
        # remember how many empty slots we have
        # non-null nodes occupy one slot but create two new slots
        # null nodes occupy one slot
        
        p = preorder.split(',')
        
        #initially we have one empty slot to put the root in it
        slot = 1
        for node in p:
            
            # no empty slot to put the current node
            if slot == 0:
                return False
                
            # a null node?
            if node == '#':
                # ocuppy slot
                slot -= 1
            else:
                # create new slot
                slot += 1
        
        #we don't allow empty slots at the end
        return slot==0
```


```python
class Solution(object):
    def isValidSerialization(self, preorder):
        """
        :type preorder: str
        :rtype: bool
        """
        if not preorder:
            return False
        
        s = []
        nodes = preorder.split(",")
        for c in nodes:
            while c == '#' and s and s[-1] == '#':
                s.pop()
                if not s:
                    return False
                s.pop()
            s.append(c)
        return len(s) == 1 and s[-1] == '#'
```