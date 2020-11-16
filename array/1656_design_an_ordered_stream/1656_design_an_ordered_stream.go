package main

type OrderedStream struct {
	nums []string
	ptr  int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{make([]string, n), 0}
}

func (this *OrderedStream) Insert(id int, value string) []string {
	id = id - 1
	this.nums[id] = value
	end := this.ptr
	for end < len(this.nums) && this.nums[end] != "" {
		end++
	}
	if end != this.ptr {
		defer func() {
			this.ptr = end
		}()
		return this.nums[this.ptr:end]
	}
	return []string{}
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(id,value);
 */
