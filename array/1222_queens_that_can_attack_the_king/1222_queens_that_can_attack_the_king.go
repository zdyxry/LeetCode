package main

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	table := make([][]int, 8)
	result := make([][]int, 0)
	for i := 0; i < 8; i++ {
		table[i] = make([]int, 8)
	}

	for i := 0; i < len(queens); i++ {
		row := queens[i][0]
		col := queens[i][1]
		table[row][col] = 1
	}

	row := king[0]
	col := king[1]
	//left
	for row < 8 {
		if table[row][col] == 1 {
			result = append(result, []int{row, col})
			break
		} else {
			row++
		}
	}
	row = king[0]
	col = king[1]
	//left
	for row >= 0 {
		if table[row][col] == 1 {
			result = append(result, []int{row, col})
			break
		} else {
			row--
		}
	}

	row = king[0]
	col = king[1]
	//left
	for col >= 0 {
		if table[row][col] == 1 {
			result = append(result, []int{row, col})
			break
		} else {
			col--
		}
	}

	row = king[0]
	col = king[1]
	//left
	for col < 8 {
		if table[row][col] == 1 {
			result = append(result, []int{row, col})
			break
		} else {
			col++
		}
	}

	row = king[0]
	col = king[1]
	//left
	for row >= 0 && col >= 0 {
		if table[row][col] == 1 {
			result = append(result, []int{row, col})
			break
		} else {
			row--
			col--
		}
	}

	row = king[0]
	col = king[1]
	//left
	for row >= 0 && col < 8 {
		if table[row][col] == 1 {
			result = append(result, []int{row, col})
			break
		} else {
			row--
			col++
		}
	}

	row = king[0]
	col = king[1]
	//left
	for row < 8 && col >= 0 {
		if table[row][col] == 1 {
			result = append(result, []int{row, col})
			break
		} else {
			row++
			col--
		}
	}

	row = king[0]
	col = king[1]
	//left
	for row < 8 && col < 8 {
		if table[row][col] == 1 {
			result = append(result, []int{row, col})
			break
		} else {
			row++
			col++
		}
	}
	return result
}
