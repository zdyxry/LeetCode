package main

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	var val, index int
	ans := []int{}
	sum := 0
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			sum += A[i]
		}
	}
	for i := 0; i < len(A); i++ {
		val = queries[i][0]
		index = queries[i][1]
		if A[index]%2 == 0 {
			sum -= A[index]
			A[index] += val
			if A[index]%2 == 0 {
				sum += A[index]
			}
		} else {
			A[index] += val
			if A[index]%2 == 0 {
				sum += A[index]
			}
		}
		ans = append(ans, sum)
	}
	return ans
}
