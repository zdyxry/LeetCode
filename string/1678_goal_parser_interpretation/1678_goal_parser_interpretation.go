package main

import "fmt"

func interpret(command string) string {
	i := 0
	ans := ""
	for i < len(command) {
		if command[i] == 'G' {
			ans += "G"
			i += 1
		} else {
			if command[i+1] == ')' {
				ans += "o"
				i += 2
			} else {
				ans += "al"
				i += 4
			}
		}
	}
	return ans
}

func main() {
	command := "G()(al)"
	res := interpret(command)
	fmt.Println(res)
}
