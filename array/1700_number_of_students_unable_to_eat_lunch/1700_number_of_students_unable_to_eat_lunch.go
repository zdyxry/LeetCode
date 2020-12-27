package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	qCount := 0
	for len(students) > 0 {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			qCount = 0
			continue
		}
		var temp []int
		temp = append(temp, students[1:]...)
		temp = append(temp, students[0])
		students = temp
		qCount++
		if qCount == len(students) {
			break
		}
	}
	return len(students)
}

func main() {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}
	res := countStudents(students, sandwiches)
	fmt.Println(res)
}
