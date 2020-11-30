package main

type FrontMiddleBackQueue struct {
	queue []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{[]int{}}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.queue = append([]int{val}, this.queue...)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	mid := len(this.queue) / 2
	this.queue = append(this.queue[:mid], append([]int{val}, this.queue[mid:]...)...)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
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
	mid := (len(this.queue) - 1) / 2
	defer func() {
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
