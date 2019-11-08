package main 

import "fmt"


func subarraysDivByK(A []int, K int) int {
	count := [10000]int{}
	count[0] = 1
	prefix, res := 0, 0
	for _, a := range A {
		prefix = (prefix + a % K + K) % K
		res += count[prefix]
		count[prefix]++
	}
	return res
}

func main() {
	A := []int{4,5,0,-2,-3,1}
	K := 5
	res := subarraysDivByK(A, K)
	fmt.Println(res)
}