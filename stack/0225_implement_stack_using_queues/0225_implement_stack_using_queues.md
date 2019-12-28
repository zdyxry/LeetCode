225. Implement Stack using Queues


Easy


Implement the following operations of a stack using queues.

push(x) -- Push element x onto stack.  
pop() -- Removes the element on top of the stack.  
top() -- Get the top element.  
empty() -- Return whether the stack is empty.  
Example:
```
MyStack stack = new MyStack();

stack.push(1);
stack.push(2);  
stack.top();   // returns 2
stack.pop();   // returns 2
stack.empty(); // returns false
```

Notes:

You must use only standard operations of a queue -- which means only push to back, peek/pop from front, size, and is empty operations are valid.

Depending on your language, queue may not be supported natively. You may simulate a queue by using a list or deque (double-ended queue), as long as you use only standard operations of a queue.

You may assume that all operations are valid (for example, no pop or top operations will be called on an empty stack).


## 方法

```go
// MyStack 是用 Queue 实现的 栈
type MyStack struct {
	a, b *Queue
}

// Constructor Initialize your data structure here.
func Constructor() MyStack {
	return MyStack{a: NewQueue(), b: NewQueue()}
}

// Push Push element x onto stack.
func (ms *MyStack) Push(x int) {
	if ms.a.Len() == 0 {
		ms.a, ms.b = ms.b, ms.a
	}
	ms.a.Push(x)
}

// Pop Removes the element on top of the stack and returns that element.
func (ms *MyStack) Pop() int {
	if ms.a.Len() == 0 {
		ms.a, ms.b = ms.b, ms.a
	}

	for ms.a.Len() > 1 {
		ms.b.Push(ms.a.Pop())
	}

	return ms.a.Pop()
}

// Top Get the top element.
func (ms *MyStack) Top() int {
	res := ms.Pop()
	ms.Push(res)
	return res
}

// Empty Returns whether the stack is empty.
func (ms *MyStack) Empty() bool {
	return (ms.a.Len() + ms.b.Len()) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

// Queue 是用于存放 int 的队列
type Queue struct {
	nums []int
}

// NewQueue 返回 *kit.Queue
func NewQueue() *Queue {
	return &Queue{nums: []int{}}
}

// Push 把 n 放入队列
func (q *Queue) Push(n int) {
	q.nums = append(q.nums, n)
}

// Pop 从 q 中取出最先进入队列的值
func (q *Queue) Pop() int {
	res := q.nums[0]
	q.nums = q.nums[1:]
	return res
}

// Len 返回 q 的长度
func (q *Queue) Len() int {
	return len(q.nums)
}

// IsEmpty 反馈 q 是否为空
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
```



```python
import collections

class MyStack(object):

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.q = collections.deque()

    def push(self, x):
        """
        Push element x onto stack.
        :type x: int
        :rtype: void
        """
        self.q.append(x)
        for i in range(len(self.q) - 1):
            self.q.append(self.q.popleft())
            

    def pop(self):
        """
        Removes the element on top of the stack and returns that element.
        :rtype: int
        """
        return self.q.popleft()
        

    def top(self):
        """
        Get the top element.
        :rtype: int
        """
        return self.q[0]

    def empty(self):
        """
        Returns whether the stack is empty.
        :rtype: bool
        """
        return not len(self.q)
```