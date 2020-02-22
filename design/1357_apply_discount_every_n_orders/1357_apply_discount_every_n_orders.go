package main

type Cashier struct {
	count         int
	discount      int
	discountCount int
	prices        map[int]int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	c := Cashier{
		count:         0,
		discount:      discount,
		discountCount: n,
		prices:        map[int]int{},
	}

	for i := 0; i < len(products); i++ {
		c.prices[products[i]] = prices[i]
	}
	return c
}

func (this *Cashier) GetBill(product []int, amount []int) (rtn float64) {
	this.count++
	total := 0
	for i := 0; i < len(product); i++ {
		total += this.prices[product[i]] * amount[i]
	}
	if this.count%this.discountCount == 0 {
		rtn = float64(total) * float64(100-this.discount) / 100
	} else {
		rtn = float64(total)
	}
	return
}
