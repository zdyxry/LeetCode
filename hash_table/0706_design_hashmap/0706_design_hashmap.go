package main

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
		p := N.next
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

func main() {

}
