package main

func construct2DArray(original []int, m int, n int) [][]int {
	k := len(original)
	if k != m*n {
		return nil
	}
	ans := make([][]int, 0, m)
	for i := 0; i < k; i += n {
		ans = append(ans, original[i:i+n])
	}
	return ans

}
