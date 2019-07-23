package main

import "fmt"

type MyQueue struct {
	q []int
}

func Constructor() MyQueue {
	return MyQueue{
		q: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	q := []int{x}
	this.q = append(q, this.q...)
}

func (this *MyQueue) Pop() int {
	v := this.q[len(this.q)-1]
	this.q = this.q[:len(this.q)-1]
	return v
}

func (this *MyQueue) Peek() int {
	return this.q[len(this.q)-1]
}

func (this *MyQueue) Empty() bool {
	if len(this.q) == 0 {
		return true
	}

	return false
}

func main() {
	myQueue := Constructor()
	myQueue.Push(1)
	myQueue.Push(2)
	fmt.Println(myQueue.Peek())
	myQueue.Pop()
	fmt.Println(myQueue.Peek())
}
