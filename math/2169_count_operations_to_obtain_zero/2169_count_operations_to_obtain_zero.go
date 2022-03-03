package main

func countOperations(num1 int, num2 int) int {
	result := 0
	for num1 > 0 && num2 > 0 {
		if num1 > num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
		result += 1
	}
	return result

}
