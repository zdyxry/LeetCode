package main

import "fmt"

type MyHashSet struct {
	arr []bool
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		arr: make([]bool, 1000000),
	}
}

func (this *MyHashSet) Add(key int) {
	this.arr[key] = true
}

func (this *MyHashSet) Remove(key int) {
	if key > len(this.arr) {
		return
	}
	this.arr[key] = false
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	if key > len(this.arr) {
		return false
	}
	return this.arr[key]
}
func main() {
	s := Constructor()
	s.Add(1)
	s.Add(2)
	s.Add(2)
	fmt.Println(s.Contains(2))
	s.Remove(2)

}
