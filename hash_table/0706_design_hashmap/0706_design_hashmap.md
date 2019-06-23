706. Design HashMap

Easy

Design a HashMap without using any built-in hash table libraries.

To be specific, your design should include these functions:

put(key, value) : Insert a (key, value) pair into the HashMap. If the value already exists in the HashMap, update the value.
get(key): Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
remove(key) : Remove the mapping for the value key if this map contains the mapping for the key.

Example:
```
MyHashMap hashMap = new MyHashMap();
hashMap.put(1, 1);          
hashMap.put(2, 2);         
hashMap.get(1);            // returns 1
hashMap.get(3);            // returns -1 (not found)
hashMap.put(2, 1);          // update the existing value
hashMap.get(2);            // returns 1 
hashMap.remove(2);          // remove the mapping for 2
hashMap.get(2);            // returns -1 (not found) 
```

Note:


All keys and values will be in the range of [0, 1000000].

The number of operations will be in the range of [1, 10000].

Please do not use the built-in HashMap library.


# 方法

与 0705 类似，可以使用一个大的list 代替，或者使用二维数组，拆分为子数组。

也可以使用链表方式实现。

碎碎念：题目写了不能使用内置 HashMap，好多靠前的答案还是使用了，不知道为了什么。。。。


```go
type MyHashMap struct {
	table []int
}

func Constructor() MyHashMap {
	return MyHashMap{
		table: make([]int, 1000001),
	}
}

func (m *MyHashMap) Put(key int, value int) {
	m.table[key] = value + 1
}

func (m *MyHashMap) Get(key int) int {
	return m.table[key] - 1
}

func (m *MyHashMap) Remove(key int) {
	m.table[key] = 0
}
```

```go
const Len int = 100000

type MyHashMap struct {
	content [Len]*Node
}

type Node struct {
	key  int
	val  int
	next *Node
}

func (N *Node) Put(key int, value int) {
	if N.key == key {
		N.val = value
		return
	}
	if N.next == nil {
		N.next = &Node{key, value, nil}
		return
	}
	N.next.Put(key, value)
}

func (N *Node) Get(key int) int {
	if N.key == key {
		return N.val
	}
	if N.next == nil {
		return -1
	}
	return N.next.Get(key)
}

func (N *Node) Remove(key int) *Node {
	if N.key == key {
        p :=N.next
        N.next = nil
        return p
	}
	if N.next != nil {
		return N.next.Remove(key)
	}
    return nil
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	node := this.content[this.Hash(key)]
	if node == nil {
		this.content[this.Hash(key)] = &Node{key, value, nil}
		return
	}
	node.Put(key, value)
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	node := this.content[this.Hash(key)]
	if node == nil {
		return -1
	}
	return node.Get(key)
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	node := this.content[this.Hash(key)]
	if node == nil {
		return
	}
	this.content[this.Hash(key)] = node.Remove(key)
}

func (this *MyHashMap) Hash(value int) int {
	return value % Len
}


```

```python
class MyHashMap(object):

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.bitmap = [[-1] * 1000 for _ in range(1001)]

    def put(self, key, value):
        """
        value will always be non-negative.
        :type key: int
        :type value: int
        :rtype: void
        """
        row, col = key / 1000, key % 1000
        self.bitmap[row][col] = value

    def get(self, key):
        """
        Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key
        :type key: int
        :rtype: int
        """
        row, col = key / 1000, key % 1000
        return self.bitmap[row][col]

    def remove(self, key):
        """
        Removes the mapping of the specified value key if this map contains a mapping for the key
        :type key: int
        :rtype: void
        """
        row, col = key / 1000, key % 1000
        self.bitmap[row][col] = -1
```


```python
class ListNode:
    def __init__(self, key, val):
        self.pair = (key, val)
        self.next = None

class MyHashMap:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.m = 1000;
        self.h = [None]*self.m
        

    def put(self, key, value):
        """
        value will always be non-negative.
        :type key: int
        :type value: int
        :rtype: void
        """
        index = key % self.m
        if self.h[index] == None:
            self.h[index] = ListNode(key, value)
        else:
            cur = self.h[index]
            while True:
                if cur.pair[0] == key:
                    cur.pair = (key, value) #update
                    return
                if cur.next == None: break
                cur = cur.next
            cur.next = ListNode(key, value)
        

    def get(self, key):
        """
        Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key
        :type key: int
        :rtype: int
        """
        index = key % self.m
        cur = self.h[index]
        while cur:
            if cur.pair[0] == key:
                return cur.pair[1]
            else:
                cur = cur.next
        return -1
            
        

    def remove(self, key):
        """
        Removes the mapping of the specified value key if this map contains a mapping for the key
        :type key: int
        :rtype: void
        """
        index = key % self.m
        cur = prev = self.h[index]
        if not cur: return
        if cur.pair[0] == key:
            self.h[index] = cur.next
        else:
            cur = cur.next
            while cur:
                if cur.pair[0] == key:
                    prev.next = cur.next
                    break
                else:
                    cur, prev = cur.next, prev.next
```
