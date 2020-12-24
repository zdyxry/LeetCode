package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	lettersMap := make(map[string]string)
	var letters []string
	var digits []string

	for _, log := range logs {
		v := strings.SplitN(log, " ", 2)
		if isDigit(v[1]) {
			digits = append(digits, log)
			continue
		} else {
			lettersMap[v[1]+v[0]] = log
		}
	}
	for k, _ := range lettersMap {
		letters = append(letters, k)
	}
	sort.Strings(letters)
	var result []string
	for _, l := range letters {
		result = append(result, lettersMap[l])
	}
	result = append(result, digits...)
	return result
}

func isDigit(log string) bool {
	for _, v := range log {
		vv := string(v)
		_, err := strconv.Atoi(vv)
		if err != nil && vv != " " {
			return false
		}
	}
	return true
}

func main() {
	logs := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
	res := reorderLogFiles(logs)
	fmt.Println(res)
}
