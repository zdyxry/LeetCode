package main

import "fmt"

func defangIPaddr(address string) string {
	var defanged []rune

	for _, char := range address {
		if char == '.' {
			defanged = append(defanged, '[', '.', ']')
		} else {
			defanged = append(defanged, char)
		}
	}
	return string(defanged)
}

func main() {
	address := "1.1.1.1"
	res := defangIPaddr(address)
	fmt.Println(res)
}
