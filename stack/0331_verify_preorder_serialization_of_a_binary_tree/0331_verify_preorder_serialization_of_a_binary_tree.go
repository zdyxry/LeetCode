package main 

import (
	"fmt"
	"strings"
)


func isValidSerialization1(preorder string) bool {
	nodes, diff := strings.Split(preorder, ","), 1
	for _, node := range nodes {
		diff--
		if diff < 0 {
			return false
		}
		if node != "#" {
			diff += 2
		}
	}
	return diff == 0
}

func main() {
	preorder := "9,3,4,#,#,1,#,#,2,#,6,#,#"
	res := isValidSerialization1(preorder)
	fmt.Println(res)
}