package main

type StockSpanner struct {
	stack [][]int
}

func Constructor() StockSpanner {
	s := StockSpanner{}
	s.stack = make([][]int, 0)

	return s
}

func (this *StockSpanner) Next(price int) int {

	ans := 1
	for len(this.stack) > 0 && price >= this.stack[len(this.stack)-1][0] {
		ans = ans + this.stack[len(this.stack)-1][1]
		this.stack = this.stack[:len(this.stack)-1]
	}

	this.stack = append(this.stack, []int{price, ans})

	return ans
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
