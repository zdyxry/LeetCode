package main

import "fmt"

// MinStack define
type MinStack struct {
	elements []int
	min      []int
	l        int
}

// Contructor define
func Contructor() MinStack {
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

// Pop define
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

func main() {
	minStack := Contructor()
	minStack.Push(1)
	minStack.Push(2)
	minStack.Push(3)
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}
