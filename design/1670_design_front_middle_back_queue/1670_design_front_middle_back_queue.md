1670. Design Front Middle Back Queue


Medium


Design a queue that supports push and pop operations in the front, middle, and back.

Implement the FrontMiddleBack class:

```
FrontMiddleBack() Initializes the queue.
void pushFront(int val) Adds val to the front of the queue.
void pushMiddle(int val) Adds val to the middle of the queue.
void pushBack(int val) Adds val to the back of the queue.
int popFront() Removes the front element of the queue and returns it. If the queue is empty, return -1.
int popMiddle() Removes the middle element of the queue and returns it. If the queue is empty, return -1.
int popBack() Removes the back element of the queue and returns it. If the queue is empty, return -1.
Notice that when there are two middle position choices, the operation is performed on the frontmost middle position choice. For example:

Pushing 6 into the middle of [1, 2, 3, 4, 5] results in [1, 2, 6, 3, 4, 5].
Popping the middle from [1, 2, 3, 4, 5, 6] returns 3 and results in [1, 2, 4, 5, 6].
```

Example 1:

```
Input:
["FrontMiddleBackQueue", "pushFront", "pushBack", "pushMiddle", "pushMiddle", "popFront", "popMiddle", "popMiddle", "popBack", "popFront"]
[[], [1], [2], [3], [4], [], [], [], [], []]
Output:
[null, null, null, null, null, 1, 3, 4, 2, -1]

Explanation:
FrontMiddleBackQueue q = new FrontMiddleBackQueue();
q.pushFront(1);   // [1]
q.pushBack(2);    // [1, 2]
q.pushMiddle(3);  // [1, 3, 2]
q.pushMiddle(4);  // [1, 4, 3, 2]
q.popFront();     // return 1 -> [4, 3, 2]
q.popMiddle();    // return 3 -> [4, 2]
q.popMiddle();    // return 4 -> [2]
q.popBack();      // return 2 -> []
q.popFront();     // return -1 -> [] (The queue is empty)
```

Constraints:

1 <= val <= 109   
At most 1000 calls will be made to pushFront, pushMiddle, pushBack, popFront, popMiddle, and popBack.


## 方法

```go
type FrontMiddleBackQueue struct {
    queue []int
}


func Constructor() FrontMiddleBackQueue {
    return FrontMiddleBackQueue{[]int{}}
}


func (this *FrontMiddleBackQueue) PushFront(val int)  {
    this.queue = append([]int{val}, this.queue...)
}


func (this *FrontMiddleBackQueue) PushMiddle(val int)  {
    mid := len(this.queue)/2
    this.queue = append(this.queue[:mid], append([]int{val}, this.queue[mid:]...)...)
}


func (this *FrontMiddleBackQueue) PushBack(val int)  {
    this.queue = append(this.queue, val)
}


func (this *FrontMiddleBackQueue) PopFront() int {
    if len(this.queue) == 0 {
        return -1
    }
    defer func() {
        this.queue = this.queue[1:]
    }()
    return this.queue[0]
}


func (this *FrontMiddleBackQueue) PopMiddle() int {
    if len(this.queue) == 0 {
        return -1
    }
    mid := (len(this.queue)-1)/2
    defer func () {
        this.queue = append(this.queue[:mid], this.queue[mid+1:]...)
    }()
    return this.queue[mid]
}


func (this *FrontMiddleBackQueue) PopBack() int {
    if len(this.queue) == 0 {
        return -1
    }
    defer func() {
        this.queue = this.queue[:len(this.queue)-1]
    }()
    return this.queue[len(this.queue)-1]
}

```

```python
class FrontMiddleBackQueue:

    def __init__(self):
        self.q = []

    def pushFront(self, val: int) -> None:
        self.q.insert(0, val)

    def pushMiddle(self, val: int) -> None:
        self.q.insert(len(self.q) // 2, val)

    def pushBack(self, val: int) -> None:
        self.q.append(val)

    def popFront(self) -> int:
        return self.q.pop(0) if self.q else -1

    def popMiddle(self) -> int:
        n = len(self.q)
        if n % 2 == 0:
            n = n // 2 - 1
        else:
            n = n // 2
        return self.q.pop(n) if self.q else -1

    def popBack(self) -> int:
        return self.q.pop() if self.q else -1


# Your FrontMiddleBackQueue object will be instantiated and called as such:
# obj = FrontMiddleBackQueue()
# obj.pushFront(val)
# obj.pushMiddle(val)
# obj.pushBack(val)
# param_4 = obj.popFront()
# param_5 = obj.popMiddle()
# param_6 = obj.popBack()
```


