705. Design HashSet

Easy

Design a HashSet without using any built-in hash table libraries.

To be specific, your design should include these functions:

add(value): Insert a value into the HashSet. 
contains(value) : Return whether the value exists in the HashSet or not.
remove(value): Remove a value in the HashSet. If the value does not exist in the HashSet, do nothing.

Example:

```
MyHashSet hashSet = new MyHashSet();
hashSet.add(1);         
hashSet.add(2);         
hashSet.contains(1);    // returns true
hashSet.contains(3);    // returns false (not found)
hashSet.add(2);          
hashSet.contains(2);    // returns true
hashSet.remove(2);          
hashSet.contains(2);    // returns false (already removed)
```

Note:


All values will be in the range of [0, 1000000].

The number of operations will be in the range of [1, 10000].

Please do not use the built-in HashSet library.


# 方法
按照题目描述，value 范围为 [0, 1000000]，可以直接用数组来实现。



```go
type MyHashSet struct {
    arr []bool
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
    return MyHashSet{
        arr : make([]bool, 1000000),
    }
}


func (this *MyHashSet) Add(key int)  {
    this.arr[key] = true
}


func (this *MyHashSet) Remove(key int)  {
    if key > len(this.arr){
        return 
    }
    this.arr[key] = false
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
    if key > len(this.arr){
        return false
    }
    return this.arr[key]
}
```


```python
class MyHashSet:
    def __init__(self):
        self.size = 10000
        self.buckets = [[] for _ in range(self.size)]

    def add(self, key):
        bucket, idx = self._index(key)
        if idx >= 0:
            return
        bucket.append(key)

    def remove(self, key):
        bucket, idx = self._index(key)
        if idx < 0:
            return
        bucket.remove(key)

    def contains(self, key):
        _, idx = self._index(key)
        return idx >= 0

    def _hash(self, key):
        return key % self.size

    def _index(self, key):
        hash = self._hash(key)
        bucket = self.buckets[hash]
        for i, k in enumerate(bucket):
            if k == key:
                return bucket, i
        return bucket, -1
```