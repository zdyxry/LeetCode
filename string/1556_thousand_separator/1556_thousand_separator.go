package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func thousandSeparator(n int) string {
	rev_s := reverse(strconv.Itoa(n))
	pattern := regexp.MustCompile(`\d{1,3}`)
	return reverse(strings.Join(pattern.FindAllString(rev_s, -1), "."))
}

func reverse(s string) string {
	temp := []rune(s)
	i := 0
	j := len(temp) - 1
	for i < j {
		temp[i], temp[j] = temp[j], temp[i]
		i++
		j--
	}
	return string(temp)
}

func main() {
	n := 987
	res := thousandSeparator(n)
	fmt.Println(res)
}
