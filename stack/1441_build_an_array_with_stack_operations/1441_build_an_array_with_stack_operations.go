package main

import "fmt"

func buildArray(target []int, n int) []string {
	var (
		ret         []string
		pushedIndex int = 0
	)

	for i := 0; i < len(target); i++ {
		if target[i] == pushedIndex+1 {
			ret = append(ret, "Push")
			pushedIndex++
		}

		if target[i] > pushedIndex+1 {
			for j := 0; j < target[i]-pushedIndex-1; j++ {
				ret = append(ret, "Push")
				ret = append(ret, "Pop")
			}

			pushedIndex += target[i] - pushedIndex - 1
			ret = append(ret, "Push")
			pushedIndex++
		}
	}
	return ret
}

func main() {
	target := []int{1, 3}
	n := 3
	res := buildArray(target, n)
	fmt.Println(res)
}
