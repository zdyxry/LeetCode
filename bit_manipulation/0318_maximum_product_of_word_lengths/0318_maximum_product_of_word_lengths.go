package main 

import "fmt"

func maxProduct(words []string) int {
	size := len(words)

	masks := make([]int, size)
	for i := 0; i < size; i++ {
		for _, b := range words[i] {
			masks[i] |= (1 << uint32(b-'a'))
		}
	}

	var max int
	for i := 0; i < size-1; i++ {
		for j := i+1; j < size; j++ {
			if masks[i]&masks[j] != 0 {
				continue
			}
			temp := len(words[i]) * len(words[j])
			if max < temp {
				max = temp
			}
		}
	}

	return max
}

func main() {
	words := []string{"abcw","baz","foo","bar","xtfn","abcdef"}
	res := maxProduct(words)
	fmt.Println(res)
}