package main

import "fmt"

func minOperations(logs []string) int {

	count := 0
	for i := range logs {
		if logs[i] == "../" {
			if count > 0 {
				count--
			}
		} else if logs[i] != "./" {
			count++
		}
	}

	return count
}

func main() {
	logs := []string{"d1/", "d2/", "../", "d21/", "./"}
	res := minOperations(logs)
	fmt.Println(res)
}
