package main

import (
	"fmt"
	"sort"
)

type item struct {
	name string
	num  int
}

type myList []item

func (list myList) Len() int {
	return len(list)
}

func (list myList) Less(i, j int) bool {
	return list[i].num > list[j].num
}

func (list myList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func longestDiverseString(a, b, c int) string {
	list := myList{item{name: "a", num: a},
		item{name: "b", num: b},
		item{name: "c", num: c}}
	sort.Sort(list)
	res := ""
	for {
		if len(res) <= 1 || res[len(res)-1] != res[len(res)-2] || string(res[len(res)-1]) != list[0].name {
			if list[0].num <= 0 {
				break
			}
			res += list[0].name
			list[0].num--
			sort.Sort(list)
		} else {
			if list[1].num <= 0 {
				break
			}
			res += list[1].name
			list[1].num--
			sort.Sort(list)
		}
	}
	return res
}

func main() {
	a := 1
	b := 1
	c := 7
	res := longestDiverseString(a, b, c)
	fmt.Println(res)
}