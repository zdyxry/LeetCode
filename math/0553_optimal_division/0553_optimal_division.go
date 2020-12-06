package main

import (
	"fmt"
	"strconv"
	"strings"
)

func optimalDivision(nums []int) string {
	var sli []string

	for _, n := range nums {
		sli = append(sli, strconv.Itoa(n))
	}

	if len(sli) < 3 {
		return strings.Join(sli, "/")
	}

	return sli[0] + "/(" + strings.Join(sli[1:], "/") + ")"
}

func main() {
	nums := []int{1000, 100, 10, 2}
	res := optimalDivision(nums)
	fmt.Println(res)
}
