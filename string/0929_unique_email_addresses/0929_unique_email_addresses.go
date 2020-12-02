package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	m := map[string]struct{}{}
	for _, email := range emails {
		i := strings.Index(email, "+")
		j := strings.LastIndex(email, "@")
		if i == -1 {
			i = j
		}
		tmp := strings.ReplaceAll(email[:i], ".", "") + email[j:]
		m[tmp] = struct{}{}
	}
	return len(m)
}

func main() {
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	res := numUniqueEmails(emails)
	fmt.Println(res)
}
