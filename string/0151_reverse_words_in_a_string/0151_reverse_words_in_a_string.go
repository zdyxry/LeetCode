package main 

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	ss := strings.Fields(s)
	reverse(&ss, 0, len(ss)-1)
	return strings.Join(ss, " ")
}

func reverse(m *[]string, i int, j int) {
	for i <= j {
		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
		i++
		j--
	}
}

func main() {
	s := "the sky is blue"
	res := reverseWords(s)
	fmt.Println(res)
}