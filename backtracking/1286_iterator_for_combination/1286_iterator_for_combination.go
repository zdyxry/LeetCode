package main

type CombinationIterator struct {
	index       int
	combination []string
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	combination := []string{}
	helper(0, combinationLength, "", characters, &combination)
	return CombinationIterator{0, combination}
}

func helper(start, length int, s, chars string, combination *[]string) {
	if length == 0 {
		*combination = append(*combination, s)
		return
	}
	for i := start; i < len(chars); i++ {
		helper(i+1, length-1, s+string(chars[i]), chars, combination)
	}
}

func (this *CombinationIterator) Next() string {
	tmp := this.combination[this.index]
	this.index++
	return tmp
}

func (this *CombinationIterator) HasNext() bool {
	return this.index < len(this.combination)
}
