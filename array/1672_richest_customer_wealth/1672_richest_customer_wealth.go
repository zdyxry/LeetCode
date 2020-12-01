package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	ret := 0
	for i := 0; i < len(accounts); i++ {
		temp := 0
		for j := 0; j < len(accounts[i]); j++ {
			temp += accounts[i][j]
		}
		if ret < temp {
			ret = temp
		}
	}
	return ret
}

func main() {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}
	res := maximumWealth(accounts)
	fmt.Println(res)
}
