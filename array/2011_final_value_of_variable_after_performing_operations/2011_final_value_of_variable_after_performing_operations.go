package main

func finalValueAfterOperations(operations []string) int {
	result := 0
	for _, o := range operations {
		if strings.Contains(o, "+") {
			result += 1
		}
		if strings.Contains(o, "-") {
			result -= 1
		}
	}
	return result

}
