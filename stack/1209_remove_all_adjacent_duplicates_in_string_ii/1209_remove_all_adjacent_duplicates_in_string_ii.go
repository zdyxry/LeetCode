package main

import "fmt"

func removeDuplicates(s string, k int) string {
	bytes := []byte(s)
	counter := make([]int, len(s))
	for i := 0; i < len(bytes); i++ {
		if i == 0 || bytes[i] != bytes[i-1] {
			counter[i] = 1
		} else {
			counter[i] = counter[i-1] + 1
			if counter[i] == k {
				bytes = append(bytes[:i-k+1], bytes[i+1:]...)
				i = i - k
			}
		}
	}
	return string(bytes)
}

func main() {
	s := "abcd"
	k := 2
	res := removeDuplicates(s, k)
	fmt.Println(res)
}
