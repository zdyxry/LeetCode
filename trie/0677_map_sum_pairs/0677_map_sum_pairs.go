package main

type MapSum struct {
	children map[rune]*MapSum
	value    int
}

func Constructor() MapSum {
	return MapSum{
		children: make(map[rune]*MapSum),
		value:    0,
	}
}

func (t *MapSum) Insert(word string, value int) {
	curr := t
	for _, ch := range word {
		if _, ok := curr.children[rune(ch)]; !ok {
			trie := Constructor()
			curr.children[rune(ch)] = &trie
		}
		curr = curr.children[rune(ch)]
	}
	curr.value = value
}

func (t *MapSum) Sum(prefix string) int {
	curr := t
	for _, ch := range prefix {
		if _, ok := curr.children[rune(ch)]; !ok {
			return 0
		}
		curr = curr.children[rune(ch)]
	}

	return SumChildren(*curr)
}

func SumChildren(t MapSum) int {
	curr := t
	sum := t.value
	for k := range curr.children {
		sum += SumChildren(*curr.children[k])
	}

	return sum
}
