package main

func minimumSum(num int) int {
	arr := make([]int, 4)
	for i := 0; i < 4; i++ {
		arr[i] = num % 10
		num /= 10
	}
	sort.Ints(arr)
	return (arr[0]+arr[1])*10 + arr[2] + arr[3]
}
