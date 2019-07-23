232. Implement Queue using Stacks


Easy


Implement the following operations of a queue using stacks.

push(x) -- Push element x to the back of queue.

pop() -- Removes the element from in front of queue.

peek() -- Get the front element.

empty() -- Return whether the queue is empty.


Example:
```

MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);  
queue.peek();  // returns 1
queue.pop();   // returns 1
queue.empty(); // returns false
```


Notes:

You must use only standard operations of a stack -- which means only push to top, peek/pop from top, size, and is empty operations are valid.

Depending on your language, stack may not be supported natively. You may simulate a stack by using a list or deque (double-ended queue), as long as you use only standard operations of a stack.

You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).


## 方法

考虑边界条件，将插入值置为最前。

```go
type MyQueue struct {
    q []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{
        q: []int{},
    }
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    q := []int{x}
    this.q = append(q, this.q...)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    v := this.q[len(this.q)-1]
    this.q = this.q[:len(this.q)-1]
    return v
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    return this.q[len(this.q)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    if len(this.q) == 0 {
        return true
    }
    return false
}
```


```python
class MyQueue(object):

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.my_queue = []

    def push(self, x):
        """
        Push element x to the back of queue.
        :type x: int
        :rtype: None
        """
        self.my_queue.insert(0, x) 

    def pop(self):
        """
        Removes the element from in front of queue and returns that element.
        :rtype: int
        """
        if self.my_queue != []:
            result = self.my_queue[-1]
            del self.my_queue[-1]
            return result
    
    def peek(self):
        """
        Get the front element.
        :rtype: int
        """
        if self.my_queue != []:
            result = self.my_queue[-1]
            #del self.my_queue[-1]
            return result
        
    def empty(self):
        """
        Returns whether the queue is empty.
        :rtype: bool
        """
        if self.my_queue == []:
            return True
        else:
            return False

```