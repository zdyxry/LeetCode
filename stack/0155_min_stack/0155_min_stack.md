155. Min Stack


Easy


Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
getMin() -- Retrieve the minimum element in the stack.
 

Example:
```
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> Returns -3.
minStack.pop();
minStack.top();      --> Returns 0.
minStack.getMin();   --> Returns -2.
```


## 方法

定义一个单独的 list，用于保存当前最小值。

```go

// MinStack define
type MinStack struct {
	elements, min []int
	l             int
}

/** initialize your data structure here. */

// Constructor155 define
func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0), 0}
}

// Push define
func (this *MinStack) Push(x int) {
	this.elements = append(this.elements, x)
	if this.l == 0 {
		this.min = append(this.min, x)
	} else {
		min := this.GetMin()
		if x < min {
			this.min = append(this.min, x)
		} else {
			this.min = append(this.min, min)
		}
	}
	this.l++
}

func (this *MinStack) Pop() {
	this.l--
	this.min = this.min[:this.l]
	this.elements = this.elements[:this.l]
}

func (this *MinStack) Top() int {
	return this.elements[this.l-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.l-1]
}
```


```python
class MinStack(object):

    def __init__(self,s=[]):
        """
        initialize your data structure here.
        """
        self.s=s
        self.t=[]
        self.m=float('inf')
        

    def push(self, x):
        """
        :type x: int
        :rtype: void
        """
        self.s.append(x)
        if x<self.m:
            self.m=x
        self.t.append(self.m)
        

    def pop(self):
        """
        :rtype: void
        """
        self.s.pop()
        self.t.pop()
        self.m=self.t[-1] if self.t else float('inf') 
        

    def top(self):
        """
        :rtype: int
        """
        return self.s[-1] if self.s else None 

    def getMin(self):
        """
        :rtype: int
        """
        return self.m if self.s else None 
```
